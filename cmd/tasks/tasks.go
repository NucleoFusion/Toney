package tasks

import "github.com/spf13/cobra"

func TasksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tasks",
		Short: "Manage your tasks",
	}

	cmd.AddCommand(
		ListCmd(),
		CreatCmd(),
		DeleteCmd(),
		EditCmd(),
	)

	return cmd
}
