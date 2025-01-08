package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kenmalik/appetizer/database"
	"github.com/kenmalik/appetizer/types"

	_ "github.com/mattn/go-sqlite3"
)

type model struct {
	msg   string
	table table.Model
  width int
  height int
}

func initialModel(t table.Model) model {
	return model{
		msg:   "Hello, World!",
		table: t,
	}
}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
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
		case "enter":
			return m, tea.Batch(
				tea.Printf("Go to %s\n", m.table.SelectedRow()[1]),
			)
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
  table := m.table.View() + "\n " + m.table.HelpView() + "\n q to quit\n"
  return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, table)
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

	var tableRows []table.Row
	for _, application := range applications {
		tableRows = append(tableRows, TableRow(application))
	}

	columns := []table.Column{
		{Title: "Company", Width: 12},
		{Title: "Position", Width: 20},
		{Title: "Location", Width: 16},
		{Title: "Posted", Width: 10},
		{Title: "Applied", Width: 10},
		{Title: "Url", Width: 16},
		{Title: "Notes", Width: 16},
		{Title: "Status", Width: 14},
	}

	style := table.DefaultStyles()
	style.Selected = style.Selected.
		Foreground(lipgloss.Color("#f5a142"))

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(tableRows),
		table.WithFocused(true),
		table.WithHeight(20),
		table.WithStyles(style),
	)

	p := tea.NewProgram(initialModel(t))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program %v", err)
		os.Exit(1)
	}
}

func TableRow(a types.Application) table.Row {
	return table.Row{
		a.Company,
		a.Position,
		a.Location,
		a.DatePosted,
		a.DateApplied,
		a.Url,
		a.Notes,
		a.Status,
	}
}
