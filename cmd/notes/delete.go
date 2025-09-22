package notes

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/spf13/cobra"
)

type DeleteOptions struct{}

func DeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Short:   "Delete a note by relative path",
		Example: "toney note delete mydir/mynote.md",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.SetConfig(); err != nil {
				return fmt.Errorf("failed to load config: %v\n\n%s", err, "Try running the `toney init` command and try again.")
			}

			if args[0] == "" {
				return fmt.Errorf("missing filepath parameter")
			}

			home, err := os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("could not get user home directory: %v", err)
			}
			path := filepath.Join(home, config.AppConfig.General.NotesDir, args[0])
			if err := os.Remove(path); err != nil {
				return fmt.Errorf("failed to delete note at %s :-\n%v", path, err)
			}

			fmt.Printf("Successfully deleted note %s", path)
			return nil
		},
	}

	return cmd
}
