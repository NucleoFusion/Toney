package keymap

import (
	"reflect"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/charmbracelet/bubbles/key"
)

type TaskPopupMap struct {
	Enter key.Binding
	Exit  key.Binding
}

func NewTaskPopupMap() TaskPopupMap {
	cfg := config.AppConfig.Keybinds.Daily
	return TaskPopupMap{
		Enter: key.NewBinding(
			key.WithKeys(cfg.Enter),
			key.WithHelp(cfg.Enter, "Enter"),
		),
		Exit: key.NewBinding(
			key.WithKeys(cfg.ExitPopup),
			key.WithHelp(cfg.ExitPopup, "Exit"),
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
	cfg := config.AppConfig.Keybinds.Daily
	return TaskFormMap{
		MoveUp: key.NewBinding(
			key.WithKeys(cfg.FormUp),
			key.WithHelp(cfg.FormUp, "Move Up"),
		),
		MoveDown: key.NewBinding(
			key.WithKeys(cfg.FormDown),
			key.WithHelp(cfg.FormDown, "Move Down"),
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
