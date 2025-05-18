package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
)

// WriteCommand represents the config:show command
var WriteCommand = &cobra.Command{
	Use:   "config:write",
	Short: "write a config file based on the current working directory",
	Long:  `Config file will be written to $HOME/.sdgo/`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Println("Config:WriteCommand called")
		if len(args) == 0 {
			return fmt.Errorf("you need to give this config a name. The name should be relevant to its usage")
		}
		configFileName := args[0]
		if configFileName == "" {
			return fmt.Errorf("it needs a name to write to the config directory so an empty string won't cut it")
		}
		// Get rid of any file extensions that may be there
		dir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		if err = os.MkdirAll(filepath.Join(dir, ".sdgo"), os.ModePerm); err != nil {
			return err
		}
		configFileName = filepath.Join(dir, ".sdgo", filepath.Base(configFileName))
		cmd.Printf("Writing config file %s \n", configFileName)
		return viper.WriteConfigAs(path.Join(dir, configFileName))

	},
}

func init() {
}
