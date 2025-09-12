package tasks

import (
	"fmt"
	"io"

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

			dl := CmdDelegate{}
			ht := dl.Height() + dl.Spacing()
			tasks := daily.TaskToItems(daily.GetItems())

			lst := list.New(tasks, dl, 1000, len(tasks)*ht)
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

type CmdDelegate struct{}

func (d CmdDelegate) Height() int                               { return 1 }
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

	switch t.Status {
	case enums.Complete:
		text += styles.CompletedStyle().Title.Render(fmt.Sprintf("%d. %s | %s", t.ID, cfg.CompletedIcon, Shorten(t.Title(), 35)))
	case enums.Pending:
		text += styles.PendingStyle().Title.Render(fmt.Sprintf("%d. %s | %s", t.ID, cfg.PendingIcon, Shorten(t.Title(), 35)))
	case enums.Started:
		text += styles.StartedStyle().Title.Render(fmt.Sprintf("%d. %s | %s", t.ID, cfg.StartedIcon, Shorten(t.Title(), 35)))
	case enums.Abandoned:
		text += styles.AbandonedStyle().Title.Render(fmt.Sprintf("%d. %s | %s", t.ID, cfg.AbandonedIcon, Shorten(t.Title(), 35)))
	}

	text = lipgloss.NewStyle().MarginLeft(3).Render(text)

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
