package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/SourcewareLab/Toney/internal/models"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "toney",
		Short: "Toney is a TUI for developers",
		Long:  `Toney is a powerful terminal UI for managing your workflows and repositories.`,
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

	cmd.AddCommand(
		ListCmd(),
		DumpCmd(),
		InitCmd(),
	)

	return cmd
}

func Execute() {
	if err := fang.Execute(context.Background(), RootCmd()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
