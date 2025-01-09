package info

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kenmalik/appetizer/types"
	"github.com/kenmalik/appetizer/view"
)


type Model struct {
  Application types.Application
}

func New(application types.Application) Model {
  return Model{
    Application: application,
  }
}

func (m Model) Update(msg tea.Msg) (view.View, tea.Cmd) {
  return m, nil
}

func (m Model) View() string {
  return m.Application.Company + " " + m.Application.Position
}
