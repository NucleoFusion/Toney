package fzf

import (
	"github.com/SourcewareLab/Toney/internal/colors"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FuzzyFinder struct {
	Width         int
	Height        int
	Vp            viewport.Model
	Ti            textinput.Model
	Items         []string
	Filtered      []string
	SelectedIndex int
}

func NewFzf(items []string, w int, h int) FuzzyFinder {
	return FuzzyFinder{
		Width:         w,
		Height:        h,
		Items:         items,
		Filtered:      items,
		Vp:            NewVP(w, h, items),
		Ti:            NewTI(w / 3),
		SelectedIndex: 0,
	}
}

func (m *FuzzyFinder) Init() tea.Cmd {
	return nil
}

func (m *FuzzyFinder) Update(msg tea.Msg) (FuzzyFinder, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "down":
			if len(m.Filtered)-1 > m.SelectedIndex {
				m.SelectedIndex += 1
			}
			if m.SelectedIndex > m.Vp.Height+m.Vp.YOffset-1 {
				m.Vp.YOffset += 1
			}
			m.UpdateVP()
			return *m, nil
		case "up":
			if 0 < m.SelectedIndex {
				m.SelectedIndex -= 1
			}
			if m.SelectedIndex < m.Vp.YOffset {
				m.Vp.YOffset -= 1
			}
			m.UpdateVP()
			return *m, nil
		}
	}

	return *m, nil
}

func (m *FuzzyFinder) View() string {
	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Center,
			lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colors.ColorPalette().FocusedBorder).Render(m.Ti.View()),
			m.Vp.View()))
}

func (m *FuzzyFinder) UpdateVP() {
	text := ""
	for k, v := range m.Filtered {
		if k == m.SelectedIndex {
			text += SelectedItemStyle(m.Width/3).Render(v) + "\n"
			continue
		}

		text += ItemLineStyle(m.Width/3).Render(v) + "\n"
	}

	m.Vp.SetContent(text)
}
