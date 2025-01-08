package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kenmalik/appetizer/database"
	"github.com/kenmalik/appetizer/types"

	_ "github.com/mattn/go-sqlite3"
)

type model struct {
	msg          string
	applications []types.Application
}

func initialModel(applications []types.Application) model {
	return model{
		msg:          "Hello, World!",
		applications: applications,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := fmt.Sprintf("%s\n\n", m.msg)

	for _, application := range m.applications {
		s += fmt.Sprintf("%v\n", application)
	}

	s += "\nPress q to quit\n"

	return s
}

type Env struct {
	applications database.ApplicationModel
}

func main() {
	initRequired := false

	_, err := os.Stat("data.db")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			initRequired = true
		} else {
			log.Fatalf("Error checking database file status - %v", err)
		}
	}

	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatalf("Error connecting to database - %v", err)
	}
	env := Env{
		applications: database.ApplicationModel{DB: db},
	}

	if initRequired {
		log.Println("No database file found. Initializing...")
		init, err := os.ReadFile("./scripts/init.sql")
		if err != nil {
			log.Fatalf("Error opening database init script - %v", err)
		}

		_, err = env.applications.DB.Exec(string(init))
		if err != nil {
			log.Fatalf("Error running database init script - %v", err)
		}
	}

	err = env.applications.InsertApplication(types.Application{
		Company:     "Some Company",
		Position:    "A position",
		Location:    "What location",
		DatePosted:  "A date",
		DateApplied: "A date",
		Url:         "fdsaf.432.edu",
		Notes:       "Funny funny yummy",
		Status:      "Rejected",
	})
  if err != nil {
    log.Fatalf("Error inserting application - %v", err)
  }

	applications, err := env.applications.All()
	if err != nil {
		log.Fatalf("Error getting applications - %v", err)
	}

	p := tea.NewProgram(initialModel(applications))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program %v", err)
		os.Exit(1)
	}
}
