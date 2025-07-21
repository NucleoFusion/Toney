package keymap

import (
	"reflect"

	"github.com/charmbracelet/bubbles/key"
)

type DiaryMap struct {
	Edit       key.Binding
	ScrollUp   key.Binding
	ScrollDown key.Binding
}

func NewDiaryMap() DiaryMap {
	return DiaryMap{
		Edit: key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "edit"),
		),
		ScrollUp: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("up", "scroll up"),
		),
		ScrollDown: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("down", "scroll down"),
		),
	}
}

func (m DiaryMap) Bindings() []key.Binding {
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
