package table

import (
	"fmt"
	"os"

	m "bsipiczki.com/jwt-go/model"
	"bsipiczki.com/jwt-go/util"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.DoubleBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Sequence(
				clearTerminalCmd(),
				tea.Printf(m.table.SelectedRow()[0]+" COPIED TO THE CLIPBOARD:\n \n%s \n", m.table.SelectedRow()[1]),
				copyToClipboardCmd(m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func copyToClipboardCmd(value string) tea.Cmd {
	return func() tea.Msg {
		util.CopyToClippboard(value)
		return nil
	}
}

func clearTerminalCmd() tea.Cmd {
	return func() tea.Msg {
		util.ClearTerminal()
		return nil
	}
}

func (m model) View() string {

	return baseStyle.Render(m.table.View()) + "\n"
}

func getLongestNameLen(inputs []m.TableInput) int {
	max := 0

	for _, input := range inputs {
		if len(input.Name) > max {
			max = len(input.Name)
		}
	}
	return max
}

func Render(inputs []m.TableInput) {
	columns := []table.Column{
		{Title: "Name", Width: getLongestNameLen(inputs)},
		{Title: "Content", Width: util.GetTermWidth() - 30},
	}

	var rows []table.Row

	for _, input := range inputs {
		rows = append(rows, table.Row{input.Name, input.Content})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(len(inputs)),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false).
		BorderLeft(true).
		BorderRight(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
