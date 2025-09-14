package tasks

import (
	"fmt"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/SourcewareLab/Toney/internal/enums"
	"github.com/SourcewareLab/Toney/internal/models/daily"
	"github.com/SourcewareLab/Toney/internal/styles"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

type CreateOptions struct {
	Title       string
	Description string
	Status      enums.TaskStatus
	Type        enums.TaskType
}

func CreatCmd() *cobra.Command {
	opts := &CreateOptions{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create a task",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.SetConfig(); err != nil {
				return fmt.Errorf("failed to load config: %v\n\n%s", err, "Try running the `toney init` command and try again.")
			}

			grps := make([]*huh.Group, 0)

			if opts.Title == "" {
				grps = append(grps, huh.NewGroup(huh.NewInput().
					Title("Title").
					Value(&opts.Title)))
			}

			if opts.Description == "" {
				grps = append(grps, huh.NewGroup(huh.NewText().
					Title("Description").
					Value(&opts.Description)))
			}

			grps = append(grps, huh.NewGroup(huh.NewSelect[enums.TaskStatus]().
				Title("Status").
				Options(
					huh.NewOption[enums.TaskStatus]("Pending", enums.Pending),
					huh.NewOption[enums.TaskStatus]("Started", enums.Started),
					huh.NewOption[enums.TaskStatus]("Abandoned", enums.Abandoned),
					huh.NewOption[enums.TaskStatus]("Complete", enums.Complete),
				).
				Value(&opts.Status)),
				huh.NewGroup(
					huh.NewSelect[enums.TaskType]().
						Title("Type").
						Options(
							huh.NewOption[enums.TaskType]("Unique", enums.UniqueTask),
							huh.NewOption[enums.TaskType]("Recurring", enums.RecurringTask),
						).
						Value(&opts.Type)))

			form := huh.NewForm(grps...).WithTheme(styles.HuhTheme())
			form.Run()

			tasks := daily.GetItems()
			tasks = append(tasks, daily.Task{
				TaskTitle: opts.Title,
				TaskType:  opts.Type,
				TaskDesc:  opts.Description,
				Status:    opts.Status,
			})

			daily.WriteItems(tasks)

			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.Title, "title", "t", "", "title for the task")
	cmd.Flags().StringVarP(&opts.Description, "desc", "d", "", "description for the task")

	return cmd
}
