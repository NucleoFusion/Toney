package config

type Config struct {
	General  GeneralConfig  `mapstructure:"general"`
	Styles   StylesConfig   `mapstructure:"styles"`
	Keybinds KeybindsConfig `mapstructure:"keybinds"`
}

type GeneralConfig struct {
	Editor      []string   `mapstructure:"editor"`
	NotesDir    string     `mapstructure:"notes_dir"`
	StartScript []string   `mapstructure:"start_script"`
	StopScript  []string   `mapstructure:"stop_script"`
	Script      []string   `mapstructure:"script"`
	Todo        TodoConfig `mapstructure:"todo"`
}

type TodoConfig struct { // TODO: Include in Docs
	Enable   bool          `mapstructure:"enable"`
	Projects []TodoProject `mapstructure:"todo_projects"`
}

type TodoProject struct {
	Name    string   `mapstructure:"name"`
	Path    string   `mapstructure:"path"`
	Exclude []string `mapstructure:"exclude"`
}
