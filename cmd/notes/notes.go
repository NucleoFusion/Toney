package notes

import "github.com/spf13/cobra"

func NotesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "note",
		Short: "Manage your notes",
	}

	cmd.AddCommand(
		ListCmd(),
		DeleteCmd(),
		CreateCmd(),
		EditCmd(),
		RenderCmd(),
	)

	return cmd
}
