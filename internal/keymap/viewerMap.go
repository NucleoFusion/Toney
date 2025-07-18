package keymap

import (
	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/charmbracelet/bubbles/key"
)

type ViewerKeyMap struct {
	ScrollUp   key.Binding
	ScrollDown key.Binding
}

func NewViewerKeyMap() ViewerKeyMap {
	cfg := config.AppConfig.Keybinds.Home
	return ViewerKeyMap{
		ScrollUp: key.NewBinding(
			key.WithKeys(cfg.ScrollUp),
			key.WithHelp("↑", "scroll up"),
		),
		ScrollDown: key.NewBinding(
			key.WithKeys(cfg.ScrollDown),
			key.WithHelp("↓", "scroll down"),
		),
	}
}
