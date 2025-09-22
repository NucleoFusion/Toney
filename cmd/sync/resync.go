package synccmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/spf13/cobra"
)

func ResyncCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resync",
		Short: "syncs notes state",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.SetConfig(); err != nil {
				return fmt.Errorf("failed to load config: %v\n\n%s", err, "Try running the `toney init` command and try again.")
			}

			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			path := filepath.Join(home, config.AppConfig.General.NotesDir)

			command := exec.Command("git", "pull", "--rebase")
			command.Dir = path

			err = command.Run()
			if err != nil {
				return err
			}

			return err
		},
	}

	return cmd
}
