package synccmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/spf13/cobra"
)

func PushCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "push",
		Short: "push to the remote url",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.SetConfig(); err != nil {
				return fmt.Errorf("failed to load config: %v\n\n%s", err, "Try running the `toney init` command and try again.")
			}

			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			path := filepath.Join(home, config.AppConfig.General.NotesDir)

			command := exec.Command("git", "push", "-u", "origin", "HEAD")
			command.Dir = path

			err = command.Run()
			if err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
