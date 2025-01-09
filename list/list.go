package list

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kenmalik/appetizer/types"
)

type Model struct {
	Table table.Model
}

func New(applications []types.Application) Model {
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

	return Model{
		Table: t,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	return m.Table.View() + "\n " + m.Table.HelpView() + "\n q to quit\n"
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			cmds = append(cmds, selectCmd(Application(m.Table.SelectedRow())))
		}
	}

	m.Table, cmd = m.Table.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

type SelectMsg types.Application

func selectCmd(application types.Application) tea.Cmd {
	return func() tea.Msg {
		return SelectMsg(application)
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

func Application(a table.Row) types.Application {
	return types.Application{
		Company:     a[0],
		Position:    a[1],
		Location:    a[2],
		DatePosted:  a[3],
		DateApplied: a[4],
		Url:         a[5],
		Notes:       a[6],
		Status:      a[7],
	}
}
