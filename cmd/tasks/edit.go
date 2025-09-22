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

type EditOptions struct {
	ID          int
	Title       string
	Description string
	Status      enums.TaskStatus
	Type        enums.TaskType
}

func EditCmd() *cobra.Command {
	opts := &EditOptions{}
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "edit a task",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.SetConfig(); err != nil {
				return fmt.Errorf("failed to load config: %v\n\n%s", err, "Try running the `toney init` command and try again.")
			}

			tasks := daily.GetItems()

			options := make([]huh.Option[int], 0)
			for k, v := range tasks {
				options = append(options, huh.NewOption(v.Title(), k))
			}

			selectform := huh.NewForm(
				huh.NewGroup(
					huh.NewSelect[int]().
						Title("Select a task to edit").
						Value(&opts.ID).
						Options(
							options...,
						),
				),
			).WithTheme(styles.HuhTheme())

			selectform.Run()

			opts.Title = tasks[opts.ID].TaskTitle
			opts.Description = tasks[opts.ID].TaskDesc
			opts.Type = tasks[opts.ID].TaskType
			opts.Status = tasks[opts.ID].Status

			statusOpts := make([]huh.Option[enums.TaskStatus], 0)
			for i := range 4 {
				opt := huh.NewOption(enums.TaskStatusMap[enums.TaskStatus(i)], enums.TaskStatus(i)).Selected(false)

				if enums.TaskStatus(i) == opts.Status {
					opt.Selected(true)
				}

				statusOpts = append(statusOpts, opt)
			}

			typeOpts := make([]huh.Option[enums.TaskType], 0)
			for i := range 2 {
				opt := huh.NewOption(enums.TaskTypeMap[enums.TaskType(i)], enums.TaskType(i)).Selected(false)

				if enums.TaskType(i) == opts.Type {
					opt.Selected(true)
				}

				typeOpts = append(typeOpts, opt)
			}

			taskform := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("Title").
						Value(&opts.Title),
				),
				huh.NewGroup(
					huh.NewText().
						Title("Description").
						Value(&opts.Description),
				),
				huh.NewGroup(huh.NewSelect[enums.TaskStatus]().
					Title("Status").
					Options(
						statusOpts...,
					).
					Value(&opts.Status)),
				huh.NewGroup(
					huh.NewSelect[enums.TaskType]().
						Title("Type").
						Options(
							typeOpts...,
						).
						Value(&opts.Type)),
			).WithTheme(styles.HuhTheme())

			taskform.Run()

			task := daily.Task{
				TaskType:  opts.Type,
				TaskTitle: opts.Title,
				TaskDesc:  opts.Description,
				Status:    opts.Status,
			}

			tasks[opts.ID] = task

			daily.WriteItems(tasks)

			return nil
		},
	}

	return cmd
}
