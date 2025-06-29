package styles

import "github.com/charmbracelet/lipgloss"

const LogoText = `
  ______                      
 /_  __/___  ____  ___  __  __
  / / / __ \/ __ \/ _ \/ / / /
 / / / /_/ / / / /  __/ /_/ / 
/_/  \____/_/ /_/\___/\__, /  
                     /____/   
`

func GetLogo(w int, h int) string {
	return lipgloss.NewStyle().Width(w).Height(h).Foreground(lipgloss.Color("#b4befe")).Align(lipgloss.Center, lipgloss.Center).Render(LogoText)
}
