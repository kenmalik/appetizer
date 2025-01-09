package info

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kenmalik/appetizer/types"
)

type Model struct {
	Application types.Application
}

func New(application types.Application) Model {
	return Model{
		Application: application,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	return m.Application.Company + " " + m.Application.Position
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}
