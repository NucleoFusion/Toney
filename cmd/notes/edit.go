package notes

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/spf13/cobra"
)

type EditOptions struct {
	Editor string
}

func EditCmd() *cobra.Command {
	opts := &EditOptions{}
	cmd := &cobra.Command{
		Use:     "edit",
		Short:   "edit a note by relative path",
		Example: "toney note edit -e 'alacritty -e bash -c hx'",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.SetConfig(); err != nil {
				return fmt.Errorf("failed to load config: %v\n\n%s", err, "Try running the `toney init` command and try again.")
			}

			if args[0] == "" {
				return fmt.Errorf("missing filepath argument")
			}

			home, err := os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("could not get user home directory: %v", err)
			}

			path := filepath.Join(home, config.AppConfig.General.NotesDir, args[0])

			editorcmd := ""
			editorargs := make([]string, 0)
			if opts.Editor == "" {
				cfg := config.AppConfig.General.Editor
				if len(cfg) < 1 {
					return fmt.Errorf("editor command in config is empty")
				}

				editorcmd = cfg[0]

				if len(cfg) >= 2 {
					editorargs = cfg[1:]
				}
			} else {
				cmdargs := strings.Split(opts.Editor, " ")

				editorcmd = cmdargs[0]
				if len(cmdargs) >= 2 {
					editorargs = cmdargs[1:]
				}
			}

			editorargs = append(editorargs, path)

			execcmd := exec.Command(editorcmd, editorargs...)

			execcmd.Stdin = os.Stdin
			execcmd.Stdout = os.Stdout
			execcmd.Stderr = os.Stderr

			if err := execcmd.Run(); err != nil {
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.Editor, "editor", "e", "", "editor to open note, will be split ' ' and then computed")

	return cmd
}
