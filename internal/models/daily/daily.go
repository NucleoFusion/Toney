package daily

import (
	"github.com/SourcewareLab/Toney/internal/enums"
	"github.com/SourcewareLab/Toney/internal/styles"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Daily struct {
	Width  int
	Height int
	List   list.Model
	Tasks  Tasks
}

func NewDaily(w int, h int) *Daily {
	items := []list.Item{
		Task{title: "Test 1", desc: "Some description", Status: enums.Complete},
		Task{title: "Test 2", desc: "Some description", Status: enums.Pending},
		Task{title: "Test 2", desc: "Some description", Status: enums.Abandoned},
		Task{title: "Test 2", desc: "Some description", Status: enums.Started},
		Task{title: "Test 2", desc: "Some description", Status: enums.Pending},
		Task{title: "Test 2", desc: "Some description", Status: enums.Pending},
		Task{title: "Test 2", desc: "Some description", Status: enums.Abandoned},
		Task{title: "Test 2", desc: "Some description", Status: enums.Abandoned},
		Task{title: "Test 2", desc: "Some description", Status: enums.Abandoned},
		Task{title: "Test 2", desc: "Some description", Status: enums.Abandoned},
		Task{title: "Test 2", desc: "Some description", Status: enums.Abandoned},
	}

	return &Daily{
		Width:  w,
		Height: h,
		List:   list.New(items, TaskDelegate{}, w/2, 2*h/3),
	}
}

func (m *Daily) Init() tea.Cmd {
	return nil
}

func (m *Daily) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	m.List, cmd = m.List.Update(msg)

	return m, cmd
}

func (m *Daily) View() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		styles.GetDailyText(m.Width, m.Height/3),
		lipgloss.Place(m.Width, 2*m.Height/3, lipgloss.Center, lipgloss.Center, m.List.View()))
}
