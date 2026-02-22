package styles

import (
	"github.com/SourcewareLab/Toney/v2/internal/colors"
	"github.com/charmbracelet/lipgloss"
)

type TaskStyle struct {
	Title lipgloss.Style
	Desc  lipgloss.Style
}

func CompletedStyle() TaskStyle {
	return TaskStyle{
		Title: lipgloss.NewStyle().Foreground(colors.ColorPalette().CompletedTask.TaskTitle),
		Desc:  lipgloss.NewStyle().Foreground(colors.ColorPalette().CompletedTask.TaskDesc),
	}
}

func StartedStyle() TaskStyle {
	return TaskStyle{
		Title: lipgloss.NewStyle().Foreground(colors.ColorPalette().StartedTask.TaskTitle),
		Desc:  lipgloss.NewStyle().Foreground(colors.ColorPalette().StartedTask.TaskDesc),
	}
}

func AbandonedStyle() TaskStyle {
	return TaskStyle{
		Title: lipgloss.NewStyle().Foreground(colors.ColorPalette().AbandonedTask.TaskTitle),
		Desc:  lipgloss.NewStyle().Foreground(colors.ColorPalette().AbandonedTask.TaskDesc),
	}
}

func PendingStyle() TaskStyle {
	return TaskStyle{
		Title: lipgloss.NewStyle().Foreground(colors.ColorPalette().PendingTask.TaskTitle),
		Desc:  lipgloss.NewStyle().Foreground(colors.ColorPalette().PendingTask.TaskDesc),
	}
}
