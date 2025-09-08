package notes

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

type RenderOptions struct {
	Raw bool
}

func RenderCmd() *cobra.Command {
	opts := &RenderOptions{}
	cmd := &cobra.Command{
		Use:     "render",
		Short:   "render a note by relative path",
		Example: "toney note render mydir/mynote --raw",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.SetConfig(); err != nil {
				return fmt.Errorf("failed to load config: %v\n\n%s", err, "Try running the `toney init` command and try again.")
			}

			if args[0] == "" {
				return fmt.Errorf("missing title argument")
			}

			home, err := os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("could not get user home directory: %v", err)
			}

			path := filepath.Join(home, config.AppConfig.General.NotesDir, args[0])

			content, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("could not create note at %s :-\n%v", path, err)
			}

			if opts.Raw {
				fmt.Print(string(content))
			} else {
				r, err := glamour.NewTermRenderer(glamour.WithStyles(config.ToGlamourStyle(config.AppConfig.Styles.Renderer)))
				if err != nil {
					return fmt.Errorf("could not create renderer from config :-\n%v", err)
				}

				fmt.Print(r.Render(string(content)))
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.Raw, "raw", "r", false, "view raw file without styles")

	return cmd
}
