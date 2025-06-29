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

const TasksText = `
    ____          _  __          ______              __        
   / __ \ ____ _ (_)/ /__  __   /_  __/____ _ _____ / /__ _____
  / / / // __  // // // / / /    / /  / __  // ___// //_// ___/
 / /_/ // /_/ // // // /_/ /    / /  / /_/ /(__  )/ ,<  (__  ) 
/_____/ \__,_//_//_/ \__, /    /_/   \__,_//____//_/|_|/____/  
                    /____/                                     `

func GetDailyText(w int, h int) string {
	return lipgloss.NewStyle().Width(w).Height(h).Foreground(lipgloss.Color("#b4befe")).
		PaddingTop(2).Align(lipgloss.Center, lipgloss.Top).Render(TasksText)
}
