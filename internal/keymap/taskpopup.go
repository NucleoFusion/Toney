package keymap

import (
	"reflect"

	"github.com/charmbracelet/bubbles/key"
)

type TaskPopupMap struct {
	Enter key.Binding
	Exit  key.Binding
}

func NewTaskPopupMap() TaskPopupMap {
	return TaskPopupMap{
		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Enter"),
		),
		Exit: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "Exit"),
		),
	}
}

func (m TaskPopupMap) Bindings() []key.Binding {
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

type TaskFormMap struct {
	MoveDown key.Binding
	MoveUp   key.Binding
}

func NewTaskFormMap() TaskFormMap {
	return TaskFormMap{
		MoveUp: key.NewBinding(
			key.WithKeys("ctrl+up"),
			key.WithHelp("ctrl+up", "Move Up"),
		),
		MoveDown: key.NewBinding(
			key.WithKeys("ctrl+down"),
			key.WithHelp("ctrl+down", "Move Down"),
		),
	}
}

func (m TaskFormMap) Bindings() []key.Binding {
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
