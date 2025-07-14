package styles

import (
	"github.com/SourcewareLab/Toney/internal/colors"
	"github.com/charmbracelet/lipgloss"
)

const LogoText = `
  ______                      
 /_  __/___  ____  ___  __  __
  / / / __ \/ __ \/ _ \/ / / /
 / / / /_/ / / / /  __/ /_/ / 
/_/  \____/_/ /_/\___/\__, /  
                     /____/   
`

func GetLogo(w int, h int) string {
	return lipgloss.NewStyle().Width(w).Height(h).Foreground(colors.ColorPalette().Lavender).Align(lipgloss.Center, lipgloss.Center).Render(LogoText)
}

const TasksText = `
    ____          _  __          ______              __        
   / __ \ ____ _ (_)/ /__  __   /_  __/____ _ _____ / /__ _____
  / / / // __ '// // // / / /    / /  / __ '// ___// //_// ___/
 / /_/ // /_/ // // // /_/ /    / /  / /_/ /(__  )/ ,<  (__  ) 
/_____/ \__,_//_//_/ \__, /    /_/   \__,_//____//_/|_|/____/  
                    /____/                                     `

func GetDailyText(w int, h int) string {
	return lipgloss.NewStyle().Width(w).Height(h).Foreground(colors.ColorPalette().Lavender).
		PaddingTop(2).Align(lipgloss.Center, lipgloss.Top).Render(TasksText)
}

const AddTasks = `
    ___        __     __   ______              __  
   /   |  ____/ /____/ /  /_  __/____ _  _____ / /__
  / /| | / __  // __  /    / /  / __ '// ___// //_/
 / ___ |/ /_/ // /_/ /    / /  / /_/ /(__  )/ ,<   
/_/  |_|\__,_/ \__,_/    /_/   \__,_//____//_/|_|  
                                                   
`

func GetAddTasks(w int, h int) string {
	return lipgloss.NewStyle().Width(w).Height(h).Foreground(colors.ColorPalette().Lavender).
		PaddingTop(2).Align(lipgloss.Center, lipgloss.Top).Render(AddTasks)
}

const SelectStatus = `
   _____        __             __     _____  __          __              
  / ___/ ___   / /___   _____ / /_   / ___/ / /_ ____ _ / /_ __  __ _____
  \__ \ / _ \ / // _ \ / ___// __/   \__ \ / __// __ '// __// / / // ___/
 ___/ //  __// //  __// /__ / /_    ___/ // /_ / /_/ // /_ / /_/ /(__  ) 
/____/ \___//_/ \___/ \___/ \__/   /____/ \__/ \__,_/ \__/ \__,_//____/  
                                                   
`

func GetSelectStatus(w int, h int) string {
	return lipgloss.NewStyle().Width(w).Height(h).Foreground(colors.ColorPalette().Lavender).
		PaddingTop(2).Align(lipgloss.Center, lipgloss.Top).Render(SelectStatus)
}
