package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/SourcewareLab/Toney/internal/models"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the Toney TUI",
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.SetConfig(); err != nil {
			log.Printf("Config error: %v\n", err)
			log.Println("Try running 'toney init' to initialize configuration")
			return
		}

		if len(config.AppConfig.General.StartScript) > 0 {
			script := strings.Join(config.AppConfig.General.StartScript, " ")
			var command *exec.Cmd

			switch runtime.GOOS {
			case "windows":
				command = exec.Command("cmd", "/C", script)
			default:
				command = exec.Command("sh", "-c", script)
			}

			command.Stdout = os.Stdout
			command.Stderr = os.Stderr

			if err := command.Run(); err != nil {
				log.Printf("Script execution error: %v\n", err)
			}
		}

		p := tea.NewProgram(models.NewRoot(), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Toney error: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
