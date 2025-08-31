package viewer

import (
	"fmt"
	"os"
	"strings"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/SourcewareLab/Toney/internal/keymap"
	"github.com/SourcewareLab/Toney/internal/messages"
	"github.com/SourcewareLab/Toney/internal/models/fzf"

	"github.com/SourcewareLab/Toney/internal/colors"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type Viewer struct {
	IsFocused   bool
	Height      int
	Width       int
	Viewport    viewport.Model
	Ready       bool
	ShowFinder  bool
	Path        string
	Content     []string
	Highlighted int
	isEditing   bool
	Renderer    *glamour.TermRenderer
	Keymap      keymap.ViewerKeyMap
	Finder      fzf.FuzzyFinder
}

func NewViewer(w int, h int) *Viewer {
	vp := viewport.New(w*3/4, h)
	vp.YOffset = 0
	vp.Style = lipgloss.NewStyle().
		Align(lipgloss.Center, lipgloss.Center).
		BorderStyle(lipgloss.RoundedBorder()).
		MarginTop(1).
		Padding(1, 1).
		BorderForeground(colors.ColorPalette().Border)
	vp.SetContent(
		lipgloss.Place(w*3/4, h-2, lipgloss.Center, lipgloss.Center,
			lipgloss.NewStyle().Foreground(colors.ColorPalette().Text).Render("Select a file to view its contents"),
		))

	r, _ := glamour.NewTermRenderer(glamour.WithStyles(config.ToGlamourStyle(config.AppConfig.Styles.Renderer)),
		glamour.WithWordWrap(w*3/4-2))

	return &Viewer{
		Viewport:  vp,
		Height:    h,
		Width:     w,
		isEditing: false,
		Keymap:    keymap.NewViewerKeyMap(),
		Renderer:  r,
	}
}

func (m Viewer) Init() tea.Cmd {
	return nil
}

func (m *Viewer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.EditorClose:
		m.Content = m.ReadFile()
		m.Viewport.SetContent(m.RenderMarkdown(m.Content, m.Width))
		return m, nil
	case messages.ChangeFileMessage:
		m.Path = msg.Path
		m.Content = m.ReadFile()
		m.Viewport.SetContent(m.RenderMarkdown(m.Content, m.Width))
		m.Viewport.YOffset = 0
		m.Highlighted = -1
		return m, nil
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

		m.Viewport.Height = msg.Height
		m.Viewport.Height = msg.Width * 3 / 4

		return m, nil
	case messages.FzfSelection:
		m.ShowFinder = false

		if !msg.Exited {
			x := -1
			for k, v := range m.Content {
				if v == msg.Selection {
					x = k
					break
				}
			}

			m.Highlighted = x
		}
		return m, nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keymap.Grep):
			m.Finder = fzf.NewFzf(m.Content, m.Width*3/4, m.Height)
			m.ShowFinder = true
			return m, nil
		}
	}

	if m.ShowFinder {
		var cmd tea.Cmd
		m.Finder, cmd = m.Finder.Update(msg)
		return m, cmd
	}

	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.Viewport, cmd = m.Viewport.Update(msg)

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Viewer) View() string {
	if m.ShowFinder {
		return m.Finder.View()
	}

	if m.IsFocused {
		m.Viewport.Style = m.Viewport.Style.BorderForeground(colors.ColorPalette().FocusedBorder)
	} else {
		m.Viewport.Style = m.Viewport.Style.BorderForeground(colors.ColorPalette().Border)
	}

	return m.Viewport.View()
}

func (m *Viewer) Header() string {
	return ""
}

func (m *Viewer) ReadFile() []string { // Change to editor type when config done
	path := strings.TrimSuffix(m.Path, "/")

	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		content = ([]byte)(fmt.Sprintf("An error occured while reading the file:%s\n%s", m.Path, err.Error()))
	}

	return strings.Split(string(content), "\n")
}

func (m *Viewer) RenderMarkdown(md []string, width int) string {
	out, _ := m.Renderer.Render(strings.Join(md, "\n"))

	if m.Highlighted != -1 {
		s := strings.Split(out, "\n")
		s[m.Highlighted] = lipgloss.NewStyle().Background(colors.ColorPalette().MenuSelectedBg).Foreground(colors.ColorPalette().MenuSelectedText).Render(md[m.Highlighted])
		out = strings.Join(s, "\n")
	}

	return out
}
