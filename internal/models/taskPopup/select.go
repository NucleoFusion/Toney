package taskpopup

import (
	"strings"

	"github.com/SourcewareLab/Toney/internal/colors"
	"github.com/SourcewareLab/Toney/internal/enums"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SelectStatus struct {
	Width    int
	Height   int
	Opts     []enums.TaskStatus
	TitleMap map[enums.TaskStatus]string
	Selected int
}

func NewSelect(w int, h int) *SelectStatus {
	return &SelectStatus{
		Width:  w,
		Height: h,
		Opts:   []enums.TaskStatus{enums.Pending, enums.Started, enums.Abandoned, enums.Complete},
		TitleMap: map[enums.TaskStatus]string{
			enums.Started:   "Started",
			enums.Pending:   "Pending",
			enums.Complete:  "Complete",
			enums.Abandoned: "Abandoned",
		},
	}
}

func (m SelectStatus) Init() tea.Cmd {
	return nil
}

func (m *SelectStatus) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "down":
			if m.Selected < len(m.Opts)-1 {
				m.Selected += 1
			}
			return m, nil
		case "up":
			if m.Selected > 0 {
				m.Selected -= 1
			}
			return m, nil
		}
	}

	return m, nil
}

func (m *SelectStatus) View() string {
	return lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.ColorPalette().Surface1).Render(m.GetText())
}

func (m SelectStatus) GetText() string {
	text := ""
	style := lipgloss.NewStyle().Width(m.Width).Padding(0, 2).Foreground(colors.ColorPalette().Lavender)

	for idx, val := range m.Opts {
		line := m.Opts[val]
		if m.Selected == idx {
			text += style.Background(colors.ColorPalette().Lavender).
				Foreground(colors.ColorPalette().Base).
				Render(m.TitleMap[line]) + "\n"

			continue
		}

		text += style.Render(m.TitleMap[line]) + "\n"
	}

	return strings.TrimSuffix(text, "\n")
}
