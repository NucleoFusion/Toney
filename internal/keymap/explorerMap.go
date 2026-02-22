package keymap

import (
	"github.com/SourcewareLab/Toney/v2/internal/config"
	"github.com/charmbracelet/bubbles/key"
)

type ExplorerKeyMap struct {
	CreateFile  key.Binding
	MoveFile    key.Binding
	RenameFile  key.Binding
	DeleteFile  key.Binding
	OpenForEdit key.Binding
	Up          key.Binding
	Down        key.Binding
}

func NewExplorerKeyMap() ExplorerKeyMap {
	cfg := config.AppConfig.Keybinds.Home

	return ExplorerKeyMap{
		CreateFile: key.NewBinding(
			key.WithKeys(cfg.Create),
			key.WithHelp(cfg.Create, "create"),
		),
		MoveFile: key.NewBinding(
			key.WithKeys(cfg.Move),
			key.WithHelp(cfg.Move, "move"),
		),
		RenameFile: key.NewBinding(
			key.WithKeys(cfg.Rename),
			key.WithHelp(cfg.Rename, "rename"),
		),
		DeleteFile: key.NewBinding(
			key.WithKeys(cfg.Delete),
			key.WithHelp(cfg.Delete, "delete"),
		),
		OpenForEdit: key.NewBinding(
			key.WithKeys(cfg.Edit),
			key.WithHelp(cfg.Edit, "edit"),
		),
		Up: key.NewBinding(
			key.WithKeys(cfg.ScrollUp),
			key.WithHelp("↑", "up"),
		),
		Down: key.NewBinding(
			key.WithKeys(cfg.ScrollDown),
			key.WithHelp("↓", "down"),
		),
	}
}
