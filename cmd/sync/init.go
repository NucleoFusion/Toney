package synccmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/spf13/cobra"
)

func InitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "initialize sync settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.SetConfig(); err != nil {
				return fmt.Errorf("failed to load config: %v\n\n%s", err, "Try running the `toney init` command and try again.")
			}

			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			path := filepath.Join(home, config.AppConfig.General.NotesDir)

			command := exec.Command("git", "init")
			command.Dir = path

			err = command.Run()
			return err
		},
	}

	return cmd
}
