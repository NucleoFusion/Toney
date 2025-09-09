package tasks

import (
	"fmt"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/SourcewareLab/Toney/internal/models/daily"
	"github.com/charmbracelet/bubbles/list"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	Search string
	All    bool
	Path   string
}

func ListCmd() *cobra.Command {
	// opts := &ListOptions{}
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all current tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.SetConfig(); err != nil {
				return fmt.Errorf("failed to load config: %v\n\n%s", err, "Try running the `toney init` command and try again.")
			}

			tasks := daily.GetItems().ItemsAsList()
			ht := daily.TaskDelegate{}.Height() + daily.TaskDelegate{}.Spacing()

			lst := list.New(tasks, daily.TaskDelegate{}, 1000, len(tasks)*ht)
			lst.SetShowTitle(false)
			lst.SetShowHelp(false)
			lst.SetShowFilter(false)
			lst.SetShowStatusBar(false)
			lst.SetShowPagination(false)

			fmt.Printf("\n\n%s\n\n", lst.View())
			return nil
		},
	}

	return cmd
}
