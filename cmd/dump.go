package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SourcewareLab/Toney/v2/internal/config"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
)

func DumpCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "dump",
		Short: "Dump the default config to config file",
		Run: func(cmd *cobra.Command, args []string) {
			def := config.DefaultConfig()

			home, _ := os.UserHomeDir()
			f, err := os.Create(filepath.Join(home, ".config", "toney", "config.toml"))
			if err != nil {
				fmt.Println("Error Creating File: ", err.Error())
				return
			}
			defer f.Close()

			encoder := toml.NewEncoder(f)
			err = encoder.Encode(def)
			if err != nil {
				fmt.Println("Failed to write to file: ", err.Error())
				return
			}

			fmt.Println("Successfully Dumped file to ~/.config/toney/config.toml")
		},
	}
}
