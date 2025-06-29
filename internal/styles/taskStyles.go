package styles

import "github.com/charmbracelet/lipgloss"

type TaskStyle struct {
	Title lipgloss.Style
	Desc  lipgloss.Style
}

var CompletedStyle = TaskStyle{
	Title: lipgloss.NewStyle().Foreground(lipgloss.Color("#a6e3a1")),
	Desc:  lipgloss.NewStyle().Foreground(lipgloss.Color("#5a7a57")),
}

var StartedStyle = TaskStyle{
	Title: lipgloss.NewStyle().Foreground(lipgloss.Color("#f9e2af")),
	Desc:  lipgloss.NewStyle().Foreground(lipgloss.Color("#a38e65")),
}

var AbandonedStyle = TaskStyle{
	Title: lipgloss.NewStyle().Foreground(lipgloss.Color("#f38ba8")),
	Desc:  lipgloss.NewStyle().Foreground(lipgloss.Color("#894454")),
}

var PendingStyle = TaskStyle{
	Title: lipgloss.NewStyle().Foreground(lipgloss.Color("#6c7086")),
	Desc:  lipgloss.NewStyle().Foreground(lipgloss.Color("#585b70")),
}
