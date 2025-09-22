package tasks

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/SourcewareLab/Toney/internal/enums"
	"github.com/SourcewareLab/Toney/internal/models/daily"
	"github.com/SourcewareLab/Toney/internal/styles"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	Raw     bool
	Verbose bool
}

func ListCmd() *cobra.Command {
	opts := &ListOptions{}
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all current tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.SetConfig(); err != nil {
				return fmt.Errorf("failed to load config: %v\n\n%s", err, "Try running the `toney init` command and try again.")
			}

			dl := CmdDelegate{Verbose: opts.Verbose}
			ht := dl.Height() + dl.Spacing()
			tasks := daily.TaskToItems(daily.GetItems())

			lst := list.New(tasks, dl, 1000, len(tasks)*ht)
			lst.SetShowTitle(false)
			lst.SetShowHelp(false)
			lst.SetShowFilter(false)
			lst.SetShowStatusBar(false)
			lst.SetShowPagination(false)

			if opts.Raw {
				path := daily.GetPath()
				content, _ := os.ReadFile(path)
				fmt.Println(string(content))
				return nil
			}
			fmt.Printf("\n\n%s\n\n", lst.View())
			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.Raw, "raw", "r", false, "get raw output of the daily tasks file")
	cmd.Flags().BoolVarP(&opts.Verbose, "verbose", "v", false, "verbose rendering of tasks, shows description as well")

	return cmd
}

type CmdDelegate struct {
	Verbose bool
}

func (d CmdDelegate) Height() int {
	if d.Verbose {
		return 2
	}

	return 1
}
func (d CmdDelegate) Spacing() int                              { return 1 }
func (d CmdDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (d CmdDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	if len(m.Items()) == 0 {
		return
	}

	t, ok := item.(daily.Task)
	if !ok {
		return
	}

	text := ""
	cfg := config.AppConfig.Styles.Icons.TaskIcons

	icon := ""
	switch t.Status {
	case enums.Complete:
		icon = cfg.CompletedIcon
	case enums.Pending:
		icon = cfg.PendingIcon
	case enums.Started:
		icon = cfg.StartedIcon
	case enums.Abandoned:
		icon = cfg.AbandonedIcon
	}

	taskText := fmt.Sprintf("%s | %s", icon, Shorten(t.Title(), 35))

	switch t.Status {
	case enums.Complete:
		taskText = styles.CompletedStyle().Title.Render(taskText)
	case enums.Pending:
		taskText = styles.PendingStyle().Title.Render(taskText)
	case enums.Started:
		taskText = styles.StartedStyle().Title.Render(taskText)
	case enums.Abandoned:
		taskText = styles.AbandonedStyle().Title.Render(taskText)
	}

	if d.Verbose {
		descText := fmt.Sprintf("\n%s      %s", strings.Repeat(" ", len(strconv.Itoa(t.ID))), t.Description()) // All the spacing to align correctly with title

		switch t.Status {
		case enums.Complete:
			descText = styles.CompletedStyle().Desc.Render(descText)
		case enums.Pending:
			descText = styles.PendingStyle().Desc.Render(descText)
		case enums.Started:
			descText = styles.StartedStyle().Desc.Render(descText)
		case enums.Abandoned:
			descText = styles.AbandonedStyle().Desc.Render(descText)
		}

		taskText += descText
	}

	text += taskText

	text = lipgloss.NewStyle().MarginLeft(3).Render(fmt.Sprintf("%d. %s", t.ID, text))

	io.WriteString(w, text)
}

func Shorten(s string, maxLen int) string {
	if len(s) <= maxLen {
		return lipgloss.NewStyle().Width(maxLen).Render(s)
	}
	if maxLen <= 1 {
		return s[:maxLen]
	}
	return s[:maxLen-1] + "…"
}
