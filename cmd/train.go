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
	trainCmd.PersistentFlags().String("training_data_dir", "text", "Directory where the training data sources should be stored at time of execution.")
	trainCmd.PersistentFlags().Int("batchSize", 1, "Batch size for the training data.")
	trainCmd.PersistentFlags().Int("imageChannels", 1, "Number of image channels.")
	trainCmd.PersistentFlags().Int("imageHeight", 255, "Image height in pixels.")
	trainCmd.PersistentFlags().Int("imageWidth", 255, "Image width in pixels.")
	trainCmd.PersistentFlags().Int("latentChannels", 1, "Number of latent channels.")
	trainCmd.PersistentFlags().Int("latentHeight", 255, "Image height in pixels.")
	trainCmd.PersistentFlags().Int("latentWidth", 255, "Image width in pixels.")
	trainCmd.PersistentFlags().Int("numEpochs", 1, "Number of epochs.")
	trainCmd.PersistentFlags().Int("maxSteps", 255, "Max number of steps.")
	trainCmd.PersistentFlags().Int("logInterval", 1, "Log interval in seconds.")
	trainCmd.PersistentFlags().Int("saveInterval", 1, "Save Interval in seconds.")
}
