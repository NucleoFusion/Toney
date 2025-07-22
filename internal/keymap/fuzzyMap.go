package keymap

import (
	"reflect"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/charmbracelet/bubbles/key"
)

type FuzzyMap struct {
	StartWriting key.Binding
	Up           key.Binding
	Down         key.Binding
	Enter        key.Binding
}

func NewFuzzyMap() FuzzyMap {
	cfg := config.AppConfig.Keybinds.Fuzz
	return FuzzyMap{
		StartWriting: key.NewBinding(
			key.WithKeys(cfg.StartWriting),
			key.WithHelp(cfg.StartWriting, "start writing"),
		),
		Up: key.NewBinding(
			key.WithKeys(cfg.Up),
			key.WithHelp(cfg.Up, "up"),
		),
		Down: key.NewBinding(
			key.WithKeys(cfg.Down),
			key.WithHelp(cfg.Down, "down"),
		),
		Enter: key.NewBinding(
			key.WithKeys(cfg.Enter),
			key.WithHelp(cfg.Enter, "enter"),
		),
	}
}

func (m FuzzyMap) Bindings() []key.Binding {
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
