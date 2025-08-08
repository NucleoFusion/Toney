package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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
			log.Fatalf("failed to load config: %v", err)
		}

		// Startup script
		script := strings.Join(config.AppConfig.General.StartScript, " ")
		command := exec.Command("bash", "-c", script)

		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err := command.Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		p := tea.NewProgram(models.NewRoot(), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Println("Alas, error")
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
