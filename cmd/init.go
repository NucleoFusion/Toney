package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/spf13/cobra"
)

func InitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize Toney configuration",
		Run: func(cmd *cobra.Command, args []string) {
			if err := config.SetConfig(); err != nil {
				home, _ := os.UserHomeDir()

				err = os.MkdirAll(home+"/.config/toney", 0o755)
				if err != nil {
					fmt.Println("Could not create config directory", err.Error())
					return
				}

				_, err = os.Create(home + "/.config/toney/config.toml")
				if err != nil {
					fmt.Println("Could not create config file", err.Error())
					return
				}
			}

			fmt.Println("Initializing Toney...")
			home, err := os.UserHomeDir()
			if err != nil {
				fmt.Println("Could not find $HOME directory", err.Error())
				return
			}

			dailyDir := filepath.Join(home, config.AppConfig.General.NotesDir, ".daily")

			err = os.MkdirAll(dailyDir, 0o755)
			if err != nil {
				fmt.Println("Could not create .toney directory", err.Error())
				return
			}

			diaryDir := filepath.Join(home, config.AppConfig.General.NotesDir, ".diary")

			err = os.MkdirAll(diaryDir, 0o755)
			if err != nil {
				fmt.Println("Could not create .diary directory", err.Error())
				return
			}

			fmt.Println("Toney setup was succesfull!")
		},
	}
}
