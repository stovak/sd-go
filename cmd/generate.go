/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io"
	"os"
	"path"

	sd "github.com/seasonjs/stable-diffusion"
	"github.com/spf13/cobra"
	"github.com/stovak/sdgo/pkg/instances"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate an image using a configured model",
	Long:  `Generate a simple cat picture with a loaded model`,
	RunE: func(cmd *cobra.Command, args []string) error {
		m, err := instances.CreateModelInstance(cmd)
		if err != nil {
			return err
		}
		cmd.Printf("Model is initialized %+v\n\n", m)
		var writers []io.Writer
		filenames := []string{
			"love_cat0.png",
		}
		outputFolder := cmd.Flag("outputs_folder").Value.String()
		cmd.Printf("Outputs folder: %s\n\n", outputFolder)
		for _, filename := range filenames {
			file, err := os.Create(path.Join(outputFolder, filename))
			if err != nil {
				return err
			}
			defer file.Close()
			writers = append(writers, file)
			cmd.Printf("Outputs file: %s\n\n", path.Join(outputFolder, filename))
		}
		return m.Predict("british short hair cat, high quality", sd.DefaultFullParams, writers)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
