package keymap

import (
	"github.com/SourcewareLab/Toney/v2/internal/config"
	"github.com/charmbracelet/bubbles/key"
)

type HomeKeyMap struct {
	FocusViewer   key.Binding
	FocusExplorer key.Binding
	BackToMenu    key.Binding
	Finder        key.Binding
}

func NewHomeKeyMap() HomeKeyMap {
	cfg := config.AppConfig.Keybinds.Home
	return HomeKeyMap{
		FocusExplorer: key.NewBinding(
			key.WithKeys(cfg.FocusExplorer),
			key.WithHelp(cfg.FocusExplorer, "file explorer"),
		),
		FocusViewer: key.NewBinding(
			key.WithKeys(cfg.FocusViewer),
			key.WithHelp(cfg.FocusViewer, "viewer"),
		),
		BackToMenu: key.NewBinding(
			key.WithKeys(cfg.BackToMenu),
			key.WithHelp(cfg.BackToMenu, "return to menu"),
		),
		Finder: key.NewBinding(
			key.WithKeys("f"),
			key.WithHelp("f", "find"),
		),
	}
}
