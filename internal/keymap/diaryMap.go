package keymap

import (
	"reflect"

	"github.com/SourcewareLab/Toney/v2/internal/config"
	"github.com/charmbracelet/bubbles/key"
)

type DiaryMap struct {
	Edit       key.Binding
	ScrollUp   key.Binding
	ScrollDown key.Binding
	OpenFinder key.Binding
	BackToMenu key.Binding
}

func NewDiaryMap() DiaryMap {
	cfg := config.AppConfig.Keybinds.Diary
	return DiaryMap{
		Edit: key.NewBinding(
			key.WithKeys(cfg.Edit),
			key.WithHelp(cfg.Edit, "edit"),
		),
		OpenFinder: key.NewBinding(
			key.WithKeys(cfg.Finder),
			key.WithHelp(cfg.Finder, "find"),
		),
		ScrollUp: key.NewBinding(
			key.WithKeys(cfg.ScrollUp),
			key.WithHelp(cfg.ScrollUp, "scroll up"),
		),
		ScrollDown: key.NewBinding(
			key.WithKeys(cfg.ScrollDown),
			key.WithHelp(cfg.ScrollDown, "scroll down"),
		),
		BackToMenu: key.NewBinding(
			key.WithKeys(cfg.BackToMenu),
			key.WithHelp(cfg.BackToMenu, "return to menu"),
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
