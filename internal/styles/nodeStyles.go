package styles

import (
	"github.com/SourcewareLab/Toney/v2/internal/colors"
	"github.com/charmbracelet/lipgloss"
)

var CurrentNodeStyle = lipgloss.NewStyle().Background(colors.ColorPalette().MenuSelectedBg).Foreground(colors.ColorPalette().MenuSelectedText)
