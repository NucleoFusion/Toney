package taskpopup

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type DeleteForm struct {
	Width      int
	Height     int
	isDeleting bool
}

func NewDeleteForm(w int, h int) *DeleteForm {
	return &DeleteForm{
		Width:      w,
		Height:     h,
		isDeleting: true,
	}
}

func (m DeleteForm) Init() tea.Cmd {
	return nil
}

func (m *DeleteForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "down":
			if m.isDeleting {
				m.isDeleting = !m.isDeleting
			}
			return m, nil
		case "up":
			if !m.isDeleting {
				m.isDeleting = !m.isDeleting
			}
			return m, nil
		}
	}

	return m, nil
}

func (m *DeleteForm) View() string {
	return lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#45475a")).Render(m.GetText())
}

func (m DeleteForm) GetText() string {
	yes := "Yes"
	no := "No"
	style := lipgloss.NewStyle().Width(m.Width).Foreground(lipgloss.Color("#b4befe"))

	if m.isDeleting {
		yes = style.Background(lipgloss.Color("#b4befe")).
			Foreground(lipgloss.Color("#1e1e2e")).
			Render(yes)
	} else {
		no = style.Background(lipgloss.Color("#b4befe")).
			Foreground(lipgloss.Color("#1e1e2e")).
			Render(no)
	}

	return fmt.Sprintf("%s\n%s", yes, no)
}
