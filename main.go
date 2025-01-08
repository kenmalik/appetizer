package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
  msg string
}

func initialModel() model {
  return model{
    msg: "Hello, World!",
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

  s += "Press q to quit\n"

  return s
}

func main() {
  p := tea.NewProgram(initialModel())
  if _, err := p.Run(); err != nil {
    fmt.Printf("Error running program %v", err)
    os.Exit(1)
  }
}