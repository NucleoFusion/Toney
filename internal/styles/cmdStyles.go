package styles

import (
	"github.com/SourcewareLab/Toney/v2/internal/colors"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func TableStyle(row, col int) lipgloss.Style {
	switch {
	case row == table.HeaderRow:
		return lipgloss.NewStyle().Align(lipgloss.Center).
			Bold(true).Italic(true).BorderForeground(colors.ColorPalette().Border).
			Foreground(lipgloss.Color("#74c7ec")).Background(lipgloss.Color("#313244"))
	default:
		return lipgloss.NewStyle().Foreground(colors.ColorPalette().Text).Padding(0, 1)
	}
}

func NewTable() *table.Table {
	t := table.New()
	t.Border(lipgloss.RoundedBorder())
	t.BorderStyle(lipgloss.NewStyle().BorderForeground(colors.ColorPalette().ErrorText))
	t.StyleFunc(TableStyle)
	return t
}
