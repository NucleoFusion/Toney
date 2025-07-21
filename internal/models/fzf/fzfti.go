package fzf

import (
	"github.com/SourcewareLab/Toney/internal/colors"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

func TextStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(colors.ColorPalette().Text)
}

func NewTI(w int) textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "Press '/' To Search..."
	ti.Prompt = "Filter : "
	ti.Width = w - 12
	ti.TextStyle = TextStyle()
	ti.PromptStyle = TextStyle()

	return ti
}
