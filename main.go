package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kenmalik/appetizer/database"
	"github.com/kenmalik/appetizer/list"
	"github.com/kenmalik/appetizer/types"
	"github.com/kenmalik/appetizer/view"

	_ "github.com/mattn/go-sqlite3"
)

type model struct {
  view view.View
  width int
  height int
}

func initialModel(applications []types.Application) tea.Model {
  return model{
    view: list.New(applications),
  }
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
  case tea.WindowSizeMsg:
    m.width = msg.Width
    m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
  m.view, cmd = m.view.Update(msg)
	return m, cmd
}

func (m model) View() string {
  page := m.view.View()
  return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, page)
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

	applications, err := env.applications.All()
	if err != nil {
		log.Fatalf("Error getting applications - %v", err)
	}

	p := tea.NewProgram(initialModel(applications), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program %v", err)
		os.Exit(1)
	}
}
