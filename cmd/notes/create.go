package notes

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/spf13/cobra"
)

type CreateOptions struct {
	Ext   string
	Title string
}

func CreateCmd() *cobra.Command {
	opts := &CreateOptions{}
	cmd := &cobra.Command{
		Use:     "create",
		Short:   "create a note by relative path",
		Example: "toney note create mydir/mynote -e txt -t 'My Note'",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.SetConfig(); err != nil {
				return fmt.Errorf("failed to load config: %v\n\n%s", err, "Try running the `toney init` command and try again.")
			}

			if args[0] == "" {
				return fmt.Errorf("missing title argument")
			}

			filename := fmt.Sprintf("%s.%s", args[0], opts.Ext)

			home, err := os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("could not get user home directory: %v", err)
			}

			path := filepath.Join(home, config.AppConfig.General.NotesDir, filename)

			f, err := os.Create(path)
			if err != nil {
				return fmt.Errorf("could not create note at %s :-\n%v", path, err)
			}

			if opts.Title != "" {
				fmt.Fprintf(f, "# %s\n", opts.Title)
			}

			fmt.Printf("Successfully created note %s\n", path)
			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.Ext, "ext", "e", "md", "file extension to use")
	cmd.Flags().StringVarP(&opts.Title, "title", "t", "", "title for the note")

	return cmd
}
