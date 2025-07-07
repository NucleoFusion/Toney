package daily

import (
	"fmt"

	"github.com/SourcewareLab/Toney/internal/enums"
	"github.com/SourcewareLab/Toney/internal/messages"
	taskpopup "github.com/SourcewareLab/Toney/internal/models/taskPopup"
	"github.com/SourcewareLab/Toney/internal/styles"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Daily struct {
	Width     int
	Height    int
	List      list.Model
	Tasks     Tasks
	Popup     *taskpopup.TaskPopup
	ShowPopup bool
}

func NewDaily(w int, h int) *Daily {
	WriteItems(Tasks{
		Recurring: []Task{
			{TaskTitle: "A", TaskDesc: "aaa", Status: enums.Abandoned},
			{TaskTitle: "A", TaskDesc: "aaa", Status: enums.Abandoned},
			{TaskTitle: "A", TaskDesc: "aaa", Status: enums.Abandoned},
		},
		Unique: []Task{
			{TaskTitle: "B", TaskDesc: "aaa", Status: enums.Abandoned},
			{TaskTitle: "B", TaskDesc: "aaa", Status: enums.Abandoned},
			{TaskTitle: "B", TaskDesc: "aaa", Status: enums.Abandoned},
		},
	})
	tasks := GetItems()

	return &Daily{
		Width:  w,
		Height: h,
		List:   list.New(tasks.ItemsAsList(), TaskDelegate{}, w/2, 2*h/3),
		Tasks:  tasks,
	}
}

func (m *Daily) Init() tea.Cmd {
	return nil
}

func (m *Daily) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.TaskPopupMessage:
	case tea.KeyMsg:
		if m.ShowPopup {
			updated, _ := m.Popup.Update(msg)
			if popup, ok := updated.(*taskpopup.TaskPopup); ok { // Type matching, cause I cant assign it straightaway
				m.Popup = popup
				return m, nil
			}
		}
		switch msg.String() {
		case "a":
			m.Popup = taskpopup.NewPopup(m.Width, m.Height, enums.Create)
			m.ShowPopup = true
			return m, nil
		case "s":
			m.Popup = taskpopup.NewPopup(m.Width, m.Height, enums.ChangeStatus)
			m.ShowPopup = true
			return m, nil
		case "d":
			m.Popup = taskpopup.NewPopup(m.Width, m.Height, enums.Delete)
			fmt.Println(m.Tasks.ItemsAsList())
			// m.ShowPopup = true
			return m, nil
		}
	}

	var cmd tea.Cmd

	m.List, cmd = m.List.Update(msg)

	return m, cmd
}

func (m *Daily) View() string {
	if m.ShowPopup {
		return m.Popup.View()
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		styles.GetDailyText(m.Width, m.Height/3),
		lipgloss.Place(m.Width, 2*m.Height/3, lipgloss.Center, lipgloss.Center, m.List.View()))
}
