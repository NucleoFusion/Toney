package synccmd

import "github.com/spf13/cobra"

func SyncCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync",
		Short: "sync your notes and tasks",
	}

	cmd.AddCommand(
		InitCmd(),
		SaveCmd(),
		RemoteCmd(),
	)

	return cmd
}
