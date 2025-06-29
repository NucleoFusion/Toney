package daily

import (
	"fmt"
	"io"

	"github.com/SourcewareLab/Toney/internal/enums"
	"github.com/SourcewareLab/Toney/internal/styles"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TaskDelegate struct{}

func (d TaskDelegate) Height() int                               { return 2 }
func (d TaskDelegate) Spacing() int                              { return 1 }
func (d TaskDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (d TaskDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	t, ok := item.(Task)
	if !ok {
		return
	}

	text := ""

	switch t.Status {
	case enums.Complete:
		text += fmt.Sprintf("%s\n%s",
			styles.CompletedStyle.Title.Render(fmt.Sprintf("%s | %s", "✓", t.Title())),
			styles.CompletedStyle.Desc.Render(t.Description()),
		)
	case enums.Pending:
		text += fmt.Sprintf("%s\n%s",
			styles.PendingStyle.Title.Render(fmt.Sprintf("%s | %s", "~", t.Title())),
			styles.PendingStyle.Desc.Render(t.Description()),
		)
	case enums.Started:
		text += fmt.Sprintf("%s\n%s",
			styles.StartedStyle.Title.Render(fmt.Sprintf("%s | %s", "○", t.Title())),
			styles.StartedStyle.Desc.Render(t.Description()),
		)
	case enums.Abandoned:
		text += fmt.Sprintf("%s\n%s",
			styles.AbandonedStyle.Title.Render(fmt.Sprintf("%s | %s", "×", t.Title())),
			styles.AbandonedStyle.Desc.Render(t.Description()),
		)
	}

	border := lipgloss.NewStyle().Border(lipgloss.NormalBorder()).
		PaddingLeft(2).
		BorderLeft(true).BorderBottom(false).BorderRight(false).BorderTop(false)

	if index == m.Index() {
		text = border.BorderForeground(lipgloss.Color("#b4befe")).Render(text)
	} else {
		text = border.BorderForeground(lipgloss.Color("#45475a")).Render(text)
	}

	io.WriteString(w, text)
}
