package synccmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/spf13/cobra"
)

func SaveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save",
		Short: "saves current notes state",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.SetConfig(); err != nil {
				return fmt.Errorf("failed to load config: %v\n\n%s", err, "Try running the `toney init` command and try again.")
			}

			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			path := filepath.Join(home, config.AppConfig.General.NotesDir)

			command := exec.Command("git", "add", ".")
			command.Dir = path

			err = command.Run()
			if err != nil {
				return err
			}

			command = exec.Command("git", "commit", "-m", fmt.Sprintf("Saved state at %s", time.Now().Format("15:04:05 02 Jan")))
			command.Dir = path

			err = command.Run()
			return err
		},
	}

	return cmd
}
