package taskpopup

import (
	"github.com/SourcewareLab/Toney/internal/styles"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Form struct {
	Width      int
	Height     int
	TitleInput textinput.Model
	DescInput  textarea.Model
}

func NewForm(w int, h int) *Form {
	tinput := textinput.New()
	tinput.Prompt = ""
	tinput.Placeholder = "Title"
	tinput.Focus()
	tinput.Width = w / 4

	ta := textarea.New()
	ta.Placeholder = "Enter Description..."
	ta.SetWidth(w / 4)

	return &Form{
		Width:      w,
		Height:     h,
		TitleInput: tinput,
		DescInput:  ta,
	}
}

func (m Form) Init() tea.Cmd {
	return nil
}

func (m *Form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+down":
			if m.TitleInput.Focused() {
				m.TitleInput.Blur()
				m.DescInput.Focus()
			}
			return m, nil
		case "ctrl+up":
			if m.DescInput.Focused() {
				m.DescInput.Blur()
				m.TitleInput.Focus()
			}
			return m, nil
		}
	}
	var cmd tea.Cmd

	if m.TitleInput.Focused() {
		m.TitleInput, cmd = m.TitleInput.Update(msg)
	} else if m.DescInput.Focused() {
		m.DescInput, cmd = m.DescInput.Update(msg)
	}

	return m, cmd
}

func (m *Form) View() string {
	formView := lipgloss.Place(m.Width, m.Height*2/3, lipgloss.Center, lipgloss.Top,
		lipgloss.JoinVertical(lipgloss.Center, m.InputView(), m.AreaView()))

	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Center, styles.GetAddTasks(m.Width, m.Height/3), formView))
}

func (m *Form) InputView() string {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#45475a")).
		Foreground(lipgloss.Color("#cdd6f4")).Padding(0, 1).
		Render("Title:" + "\n" + m.TitleInput.View())
}

func (m *Form) AreaView() string {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#45475a")).
		Foreground(lipgloss.Color("#cdd6f4")).Padding(0, 1).PaddingLeft(2).
		Render(m.DescInput.View())
}
