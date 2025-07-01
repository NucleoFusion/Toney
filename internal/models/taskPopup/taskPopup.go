package taskpopup

import (
	"github.com/SourcewareLab/Toney/internal/enums"
	"github.com/SourcewareLab/Toney/internal/styles"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TaskPopup struct {
	Width  int
	Height int
	Type   enums.TaskPopup
	Form   *Form
	Select *SelectStatus
	Title  string
	Desc   string
	Status enums.TaskStatus
}

func NewPopup(w int, h int, typ enums.TaskPopup) *TaskPopup {
	return &TaskPopup{
		Width:  w,
		Height: h,
		Type:   typ,
		Form:   NewForm(w, h),
		Select: NewSelect(w/3, h),
	}
}

func (m *TaskPopup) Init() tea.Cmd {
	return nil
}

func (m *TaskPopup) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.Type {
	case enums.Create:
		updated, _ := m.Form.Update(msg)
		if form, ok := updated.(*Form); ok { // Type matching, cause I cant assign it straightaway
			m.Form = form
			return m, nil
		}
	case enums.ChangeStatus:
		updated, _ := m.Select.Update(msg)
		if sel, ok := updated.(*SelectStatus); ok { // Type matching, cause I cant assign it straightaway
			m.Select = sel
			return m, nil
		}
	}

	return m, nil
}

func (m *TaskPopup) View() string {
	view := ""

	switch m.Type {
	case enums.Create:
		view = m.Form.View()
	case enums.ChangeStatus:
		view = lipgloss.JoinVertical(lipgloss.Center,
			styles.GetSelectStatus(m.Width, m.Height/2),
			lipgloss.Place(m.Width, m.Height/2, lipgloss.Center, lipgloss.Top, m.Select.View()))
	}
	return view
}
