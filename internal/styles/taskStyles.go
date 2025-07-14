package styles

import (
	"github.com/SourcewareLab/Toney/internal/colors"
	"github.com/charmbracelet/lipgloss"
)

type TaskStyle struct {
	Title lipgloss.Style
	Desc  lipgloss.Style
}

var CompletedStyle = TaskStyle{
	Title: lipgloss.NewStyle().Foreground(colors.ColorPalette().Green),
	Desc:  lipgloss.NewStyle().Foreground(colors.ColorPalette().GreenDull),
}

var StartedStyle = TaskStyle{
	Title: lipgloss.NewStyle().Foreground(colors.ColorPalette().Yellow),
	Desc:  lipgloss.NewStyle().Foreground(colors.ColorPalette().YellowDull),
}

var AbandonedStyle = TaskStyle{
	Title: lipgloss.NewStyle().Foreground(colors.ColorPalette().Red),
	Desc:  lipgloss.NewStyle().Foreground(colors.ColorPalette().RedDull),
}

var PendingStyle = TaskStyle{
	Title: lipgloss.NewStyle().Foreground(colors.ColorPalette().Overlay0),
	Desc:  lipgloss.NewStyle().Foreground(colors.ColorPalette().Surface0),
}
