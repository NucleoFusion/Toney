package styles

import (
	"github.com/SourcewareLab/Toney/v2/internal/colors"
	"github.com/charmbracelet/lipgloss"
)

func BorderStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Padding(1, 2).
		Align(lipgloss.Center, lipgloss.Center).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(colors.ColorPalette().Border)
}
