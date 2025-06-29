package menu

import (
	"strings"

	"github.com/SourcewareLab/Toney/internal/enums"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MenuList struct {
	Width      int
	Height     int
	Options    map[enums.Page]string
	Selections []enums.Page
	Selected   int
}

func NewMenuList(w int, h int, opts map[enums.Page]string) *MenuList {
	selections := make([]enums.Page, 0, len(opts))

	for k := range opts {
		selections = append(selections, k)
	}

	return &MenuList{
		Width:      w,
		Height:     h,
		Options:    opts,
		Selections: selections,
		Selected:   0,
	}
}

func (m *MenuList) Init() tea.Cmd {
	return nil
}

func (m *MenuList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "down":
			if m.Selected < len(m.Selections)-1 {
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

func (m *MenuList) View() string {
	return lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#45475a")).Render(m.GetText())
}

func (m *MenuList) GetText() string {
	text := ""
	style := lipgloss.NewStyle().Width(m.Width).Padding(0, 2).Foreground(lipgloss.Color("#b4befe"))

	for idx, val := range m.Selections {
		line := m.Options[val]
		if m.Selected == idx {
			text += style.Background(lipgloss.Color("#b4befe")).
				Foreground(lipgloss.Color("#1e1e2e")).
				Render(line) + "\n"

			continue
		}

		text += style.Render(line) + "\n"
	}

	return strings.TrimSuffix(text, "\n")
}
