package daily

import (
	"github.com/SourcewareLab/Toney/internal/enums"
	"github.com/charmbracelet/bubbles/list"
)

type Tasks struct {
	Recurring []list.Item `json:"recurring"`
	Unique    []list.Item `json:"unique"`
}

type Task struct {
	title, desc string
	Status      enums.TaskStatus
}

func (m Task) Title() string       { return m.title }
func (m Task) Description() string { return m.desc }
func (m Task) FilterValue() string { return m.title }
