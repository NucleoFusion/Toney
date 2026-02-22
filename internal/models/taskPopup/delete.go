package taskpopup

import (
	"fmt"

	"github.com/SourcewareLab/Toney/v2/internal/colors"
	"github.com/SourcewareLab/Toney/v2/internal/config"
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
		case config.AppConfig.Keybinds.Global.Down:
			if m.isDeleting {
				m.isDeleting = !m.isDeleting
			}
			return m, nil
		case config.AppConfig.Keybinds.Global.Up:
			if !m.isDeleting {
				m.isDeleting = !m.isDeleting
			}
			return m, nil
		}
	}

	return m, nil
}

func (m *DeleteForm) View() string {
	return lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.ColorPalette().Border).Render(m.GetText())
}

func (m DeleteForm) GetText() string {
	yes := "Yes"
	no := "No"
	style := lipgloss.NewStyle().Width(m.Width).Foreground(colors.ColorPalette().Text)

	if m.isDeleting {
		yes = style.Background(colors.ColorPalette().MenuSelectedBg).
			Foreground(colors.ColorPalette().MenuSelectedText).
			Render(yes)
		no = style.Render(no)
	} else {
		no = style.Background(colors.ColorPalette().MenuSelectedBg).
			Foreground(colors.ColorPalette().MenuSelectedText).
			Render(no)
		yes = style.Render(yes)
	}

	return fmt.Sprintf("%s\n%s", yes, no)
}
