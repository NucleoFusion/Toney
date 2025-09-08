package notes

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/SourcewareLab/Toney/internal/styles"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	Search string
	All    bool
	Path   string
}

func ListCmd() *cobra.Command {
	opts := &ListOptions{}
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all notes in the notes directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.SetConfig(); err != nil {
				return fmt.Errorf("failed to load config: %v\n\n%s", err, "Try running the `toney init` command and try again.")
			}

			home, _ := os.UserHomeDir()
			rootpath := filepath.Join(home, config.AppConfig.General.NotesDir)
			entries := map[string]os.DirEntry{}

			if opts.Path != "" {
				rootpath = filepath.Join(rootpath, opts.Path)
			}

			filepath.WalkDir(rootpath, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					if opts.Path == "" {
						log.Fatalf("failed to read notes directory (%s) : %v", path, err)
					} else {
						log.Fatalf("failed to read directory given with -p/--path flag (%s) : %v", path, err)
					}
				}

				if d.IsDir() {
					if d.Name() == config.AppConfig.General.NotesDir {
						return nil
					} else if strings.HasPrefix(d.Name(), ".") && !opts.All {
						return fs.SkipDir
					}
				}

				if !d.IsDir() {
					entries[path] = d
				}

				return nil
			})

			filteredEntries := entries
			// Fuzzy Filtering
			if opts.Search != "" {
				filteredEntries = map[string]os.DirEntry{} // Resetting filteredEntries

				paths := make([]string, 0, len(entries))
				for k := range entries {
					paths = append(paths, k)
				}

				res := fuzzy.Find(opts.Search, paths)

				for _, v := range res {
					filteredEntries[v] = entries[v]
				}
			}

			content := FormatEntries(filteredEntries, rootpath)

			fmt.Println(content)
			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.Search, "search", "s", "", "fuzzy find a filename that matches the given input")
	cmd.Flags().BoolVarP(&opts.All, "all", "a", false, "include hidden files in list")
	cmd.Flags().StringVarP(&opts.Path, "path", "p", "", "list files in given path, path is relative to the `notes_dir`")

	return cmd
}

func FormatEntries(entries map[string]os.DirEntry, root string) string {
	t := styles.NewTable()
	t.Headers("Filename", "Updated", "Size")

	for k, v := range entries {
		info, _ := v.Info()
		p, _ := filepath.Rel(root, k)
		t.Row(p, info.ModTime().Format("02 Jan"), fmt.Sprintf("%dKb", info.Size()))
	}

	return t.Render()
}
