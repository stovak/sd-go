package config

import (
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var ConfigCheckCmd = &cobra.Command{
	Use:   "config:check",
	Short: "Check Current Config",
	RunE:  ConfigCheck,
}

func ConfigCheck(cmd *cobra.Command, args []string) error {
	cmd.Println("Config Check")
	// 0. Show root path
	rootPath, err := cmd.Flags().GetString("root_path")
	if err != nil {
		return err
	}
	cmd.Printf("✅ Root Path: %s \n", rootPath)
	// 1. Check to see if the model file exists
	modelPath, err := cmd.Flags().GetString("model_path")
	if err != nil {
		return err
	}
	if _, err := os.Stat(modelPath); err != nil {
		return err
	}
	cmd.Printf("✅ Model Path: %s \n", modelPath)

	checkpointFile, err := cmd.Flags().GetString("checkpoint_filename")
	if err != nil {
		return err
	}
	if _, err := os.Stat(filepath.Join(modelPath, checkpointFile+".ckpt")); err != nil {
		return err
	}
	cmd.Printf("✅ Checkpoint File: %s \n", checkpointFile)

	outputsFolder, err := cmd.Flags().GetString("outputs_folder")
	if err != nil {
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(filepath.Join(rootPath, outputsFolder)); err != nil {
		cmd.Println("Outputs folder does not exist. Attempting to create.")
		cobra.CheckErr(os.MkdirAll(outputsFolder, os.ModePerm))
	}
	// 2. Check to see if the outputs folder exists
	cmd.Printf("✅ Outputs Folder: %s \n", filepath.Join(rootPath, outputsFolder))
	// Put other config checks here
	return nil
}
