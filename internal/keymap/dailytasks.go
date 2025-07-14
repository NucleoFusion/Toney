package keymap

import (
	"reflect"

	"github.com/charmbracelet/bubbles/key"
)

type DailyTaskMap struct {
	CreateTask   key.Binding
	EditTask     key.Binding
	ChangeStatus key.Binding
	DeleteTask   key.Binding
	BackToMenu   key.Binding
}

func NewDailyTaskMap() DailyTaskMap {
	return DailyTaskMap{
		CreateTask: key.NewBinding(
			key.WithKeys("c"),
			key.WithHelp("c", "create"),
		),
		EditTask: key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "edit"),
		),
		ChangeStatus: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "status change"),
		),
		DeleteTask: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "delete"),
		),
		BackToMenu: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "return to menu"),
		),
	}
}

func (m DailyTaskMap) Bindings() []key.Binding {
	var bindings []key.Binding

	v := reflect.ValueOf(m)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if binding, ok := field.Interface().(key.Binding); ok {
			bindings = append(bindings, binding)
		}
	}

	return bindings
}
