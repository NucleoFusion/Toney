package config

type Config struct {
	General  GeneralConfig  `mapstructure:"general"`
	Styles   StylesConfig   `mapstructure:"styles"`
	Keybinds KeybindsConfig `mapstructure:"keybinds"`
}

type GeneralConfig struct {
	Editor   string `mapstructure:"editor"`
	NotesDir string `mapstructure:"notes_dir"`
}
