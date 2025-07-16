package config

import "github.com/spf13/viper"

func SetConfig() {
	viper.SetConfigFile("config")
	viper.SetConfigFile("toml")
	viper.AddConfigPath("$HOME/.config/toney")
}

func DefaultConfig() Config {
	return Config{
		General: GeneralConfig{
			Editor:   "nvim",
			NotesDir: ".toney", // From $HOME Directory
		},
		Keybinds: KeybindsConfig{
			Home: HomeKeybinds{
				FocusViewer:   "V",
				FocusExplorer: "F",
				Create:        "c",
				Rename:        "r",
				ScrollUp:      "up",
				ScrollDown:    "down",
				Move:          "m",
				Delete:        "d",
				Edit:          "enter",
			},
			Daily: DailyKeybinds{
				Create:       "c",
				Edit:         "e",
				StatusChange: "s",
				Delete:       "d",
				ExitPopup:    "esc",
				Enter:        "enter",
				FormUp:       "ctrl+up",
				FormDown:     "ctrl+down",
				BackToMenu:   "esc",
			},
		},
		Styles: StylesConfig{
			Text:          "#cdd6f4",
			Background:    "#1e1e2e",
			Border:        "#45475a",
			FocusedBorder: "#b4befe",
			Icons: IconsConfig{
				FolderIcon: "📁",
				FileIcon:   "📄",
				TaskIcons: TaskIcons{
					CompletedIcon: "✓",
					AbandonedIcon: "×",
					PendingIcon:   "~",
					StartedIcon:   "○",
				},
			},
			TaskStyles: TaskStylesConfig{
				FocusedBar:   "#b4befe",
				UnfocusedBar: "#45475a",
				// TODO: Fill this
				CompletedStyle: TaskStateStyle{},
				AbandonedStyle: TaskStateStyle{},
				PendingStyle:   TaskStateStyle{},
				StartedStyle:   TaskStateStyle{},
			},
		},
	}
}
