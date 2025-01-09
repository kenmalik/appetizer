package info

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kenmalik/appetizer/types"
)

var (
	headerStyles = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true)
	linkStyles = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f5a142")).
      Underline(true)
)

type Model struct {
	application types.Application
}

func New(application types.Application) Model {
	return Model{
		application: application,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	s := headerStyles.Render(fmt.Sprintf("   %s - %s   ", m.application.Company, m.application.Position)) + "\n\n"

	s += m.application.Location + "\n\n"

	s += "Posted: " + m.application.DatePosted + "\n\n"

	s += "Source: " + linkStyles.Render(m.application.Url) + "\n"

  s += "\nStatus: " + m.application.Status + "\n"
	s += "Applied: " + m.application.DateApplied + "\n"

	s += "\nAdditional Notes:\n" + m.application.Notes

	return s
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}
