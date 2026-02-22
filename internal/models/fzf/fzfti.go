package fzf

import (
	"fmt"

	"github.com/SourcewareLab/Toney/v2/internal/colors"
	"github.com/SourcewareLab/Toney/v2/internal/config"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

func TextStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(colors.ColorPalette().Text)
}

func NewTI(w int) textinput.Model {
	ti := textinput.New()
	ti.Placeholder = fmt.Sprintf("Press '%s' To Search...", config.AppConfig.Keybinds.Fuzz.StartWriting)
	ti.Prompt = "Filter : "
	ti.Width = w - 12
	ti.TextStyle = TextStyle()
	ti.PromptStyle = TextStyle()

	return ti
}
