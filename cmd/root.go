/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stovak/sdgo/cmd/config"
	"os"
	"path"
)

var (
	cfgFile string
	debug   bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sdgo",
	Short: "Command Line Interface for stable diffusion written in go",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Print(`
***COMMAND CONFIG VALUE PRECEDENCE***: 

	DEFAULT 
	  	is overwritten by
	CONFIG_FILE
		is overwritten by 
	SDGO_<ENV_VAR>
		(e.g. export SDGO_ROOT_PATH=${HOME})

`)
		return config.ConfigCheck(cmd, args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	err := initConfig(rootCmd)
	if err != nil {
		cobra.CheckErr(err)
	}
	dir, err := os.Getwd()
	if err != nil {
		cobra.CheckErr(err)
	}
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sdgo/sdgo.yaml)")
	rootCmd.PersistentFlags().String("root_path", dir, "Always the CWD unless set with this flag/ENV_VAR")
	rootCmd.PersistentFlags().String("model_path", "models", "A help for foo. always relative to ROOT_PATH")
	rootCmd.PersistentFlags().String("checkpoint_filename", "v1-5-pruned-emaonly", "The filename of the checkpoint to be used as the model")
	rootCmd.PersistentFlags().String("outputs_folder", "output", "folder to write all outputs. if it doesn't exist, it will be created.")
	rootCmd.PersistentFlags().String("log_level", "info", "Log level")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", true, "verbose logging")

	rootCmd.AddCommand(trainCmd)
	rootCmd.AddCommand(config.ShowCmd)
	rootCmd.AddCommand(config.WriteCommand)
}

// initConfig reads in config file and ENV variables if set.
func initConfig(cmd *cobra.Command) error {
	cmd.Printf("Loading config file... %s \n", cfgFile)
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
	var err error
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".sdgo" (without extension).
		viper.AddConfigPath(path.Join(home, ".sdgo"))
		viper.SetConfigType("yaml")
		viper.SetConfigName("sdgo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err = viper.ReadInConfig(); err == nil {
		log.Debugf("Using config file: %s \n", viper.ConfigFileUsed())
		return nil
	}
	cmd.Println("No config file found, using defaults from config.dist.yaml")
	// read in from the defaults config
	f, err := os.Open("config.dist.yaml")
	cobra.CheckErr(err)
	log.Debug("Using config file: config.dist.yaml")
	return viper.ReadConfig(f)
}

type PlainFormatter struct {
}

func (f *PlainFormatter) Format(entry *log.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s\n", entry.Message)), nil
}
func toggleDebug(cmd *cobra.Command, args []string) {
	if debug {
		log.Info("Debug logs enabled")
		log.SetLevel(log.DebugLevel)
		log.SetFormatter(&log.TextFormatter{})
	} else {
		plainFormatter := new(PlainFormatter)
		log.SetFormatter(plainFormatter)
	}

}
