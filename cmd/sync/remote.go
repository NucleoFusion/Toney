package synccmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/spf13/cobra"
)

func RemoteCmd() *cobra.Command {
	var IsSetting bool
	cmd := &cobra.Command{
		Use:   "remote",
		Short: "set the remote url",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.SetConfig(); err != nil {
				return fmt.Errorf("failed to load config: %v\n\n%s", err, "Try running the `toney init` command and try again.")
			}

			if args[0] == "" {
				return fmt.Errorf("Empty remote url found")
			}

			remote := args[0]

			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			path := filepath.Join(home, config.AppConfig.General.NotesDir)

			if IsSetting {
				command := exec.Command("git", "remote", "set-url", "origin", remote)
				command.Dir = path

				err = command.Run()
				if err != nil {
					return err
				}
			} else {
				command := exec.Command("git", "remote", "add", "origin", remote)
				command.Dir = path

				err = command.Run()
				if err != nil {
					return err
				}
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&IsSetting, "set", "s", false, "set one you are changing the remote url")

	return cmd
}
