package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/SourcewareLab/Toney/internal/styles"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	Search string
}

func ListCmd() *cobra.Command {
	opts := &ListOptions{}
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all notes in the notes directory",
		Run: func(cmd *cobra.Command, args []string) {
			if err := config.SetConfig(); err != nil {
				log.Fatalf("failed to load config: %v\n%s", err, "Try running the `toney init` command and try again.")
			}

			home, _ := os.UserHomeDir()
			rootpath := filepath.Join(home, config.AppConfig.General.NotesDir)
			entries := map[string]os.DirEntry{}

			filepath.WalkDir(rootpath, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					log.Fatalf("failed to read notes directory (%s) : %v", path, err)
				}

				if d.IsDir() {
					if strings.HasPrefix(d.Name(), ".") && d.Name() != config.AppConfig.General.NotesDir {
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
		},
	}

	cmd.Flags().StringVarP(&opts.Search, "search", "s", "", "fuzzy find a filename that matches the given input")

	return cmd
}

func FormatEntries(entries map[string]os.DirEntry, root string) string {
	t := styles.NewTable()
	t.Headers("Filename", "Updated", "Size")

	for k, v := range entries {
		info, _ := v.Info()
		p, _ := filepath.Rel(root, k)
		t.Row(p, info.ModTime().Format(time.DateTime), fmt.Sprintf("%dKb", info.Size()))
	}

	return t.Render()
}
