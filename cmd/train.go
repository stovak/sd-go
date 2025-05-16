package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stovak/sdgo/pkg/instances"
)

// trainCmd represents the train command
var trainCmd = &cobra.Command{
	Use:   "train",
	Short: "Train Stable diffusion on a folder of character images",
	Long:  `The command will look through the config file for `,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := instances.CreateModelInstance(cmd)
		if err != nil {
			return err
		}
		cmd.Println("Model is initialized")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(trainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	rootCmd.PersistentFlags().String("model_path", "models", "A help for foo")
	rootCmd.PersistentFlags().String("checkpoint_filename", "v2-1_768-ema-pruned", "A help for foo")
	// trainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
