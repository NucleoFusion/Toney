package tasks

import (
	"fmt"
	"slices"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/SourcewareLab/Toney/internal/models/daily"
	"github.com/SourcewareLab/Toney/internal/styles"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

type DeleteOptions struct {
	ID      int
	Confirm bool
}

func DeleteCmd() *cobra.Command {
	opts := &DeleteOptions{}
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete a task",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.SetConfig(); err != nil {
				return fmt.Errorf("failed to load config: %v\n\n%s", err, "Try running the `toney init` command and try again.")
			}

			tasks := daily.GetItems()

			options := make([]huh.Option[int], 0)
			for _, v := range tasks {
				options = append(options, huh.NewOption[int](v.Title(), v.ID))
			}

			form := huh.NewForm(
				huh.NewGroup(
					huh.NewSelect[int]().
						Title("Choose task to delete").
						Value(&opts.ID).
						Options(
							options...,
						),
				),
				huh.NewGroup(
					huh.NewConfirm().
						Title(fmt.Sprintf("Delete Task: `%s` ?", tasks[opts.ID].Title())).
						Value(&opts.Confirm),
				),
			).WithTheme(styles.HuhTheme())

			form.Run()

			if !opts.Confirm {
				return nil
			}

			fmt.Printf("Deleted Task: `%s`\n", tasks[opts.ID].Title())
			tasks = slices.Delete(tasks, opts.ID, opts.ID+1)

			daily.WriteItems(tasks)

			return nil
		},
	}

	return cmd
}
