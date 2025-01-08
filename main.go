package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kenmalik/appetizer/database"

	_ "github.com/mattn/go-sqlite3"
)

type model struct {
  msg string
  applications []database.Application
}

func initialModel(applications []database.Application) model {
  return model{
    msg: "Hello, World!",
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
  db, err := sql.Open("sqlite3", "data.db")
  if err != nil {
    log.Fatalf("Error connecting to database - %v", err)
  }

  env := Env{
    applications: database.ApplicationModel{DB: db},
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
