package config

import (
	"github.com/spf13/cobra"
	"github.com/stovak/sdgo/pkg/training"
	"strconv"
)

type TrainConfig struct {
	// Now she's back in my atmosphere
	// with drops of jupiter in her hair, heh-heh

	// === Model Input Shapes ===
	BatchSize      int `yaml:"batch_size"`      // e.g., 1, 2, 4
	ImageChannels  int `yaml:"image_channels"`  // typically 3 for RGB
	ImageHeight    int `yaml:"image_height"`    // e.g., 512
	ImageWidth     int `yaml:"image_width"`     // e.g., 512
	LatentChannels int `yaml:"latent_channels"` // typically 4 (for SD 1.5)
	LatentHeight   int `yaml:"latent_height"`   // usually ImageHeight / 8 (e.g., 64)
	LatentWidth    int `yaml:"latent_width"`    // usually ImageWidth / 8 (e.g., 64)
	NumEpochs      int `yaml:"num_epochs"`      // e.g., 10
	MaxSteps       int `yaml:"max_steps"`       // total training steps (overrides epochs if set)
	LogInterval    int `yaml:"log_interval"`    // log every N steps
	SaveInterval   int `yaml:"save_interval"`   // save every N steps

	// === File System / Paths ===
	OutputDir string `yaml:"output_dir"` // where to write logs/models

	ModelPath string `yaml:"model_path"` // path to .ckpt or .safetensors if used

	TrainingDataDir string `yaml:"training_data_dir"`

	CommandStepLogger *CommandStepLogger
}

func NewTrainConfig(cmd *cobra.Command) *TrainConfig {
	batchSize, _ := strconv.Atoi(cmd.Flag("batch_size").Value.String())
	imageChannels, _ := strconv.Atoi(cmd.Flag("image_channels").Value.String())
	imageHeight, _ := strconv.Atoi(cmd.Flag("image_height").Value.String())
	imageWidth, _ := strconv.Atoi(cmd.Flag("image_width").Value.String())
	latentChannels, _ := strconv.Atoi(cmd.Flag("latent_channels").Value.String())
	latentHeight, _ := strconv.Atoi(cmd.Flag("latent_height").Value.String())
	latentWidth, _ := strconv.Atoi(cmd.Flag("latent_width").Value.String())
	numEpochs, _ := strconv.Atoi(cmd.Flag("num_epochs").Value.String())
	maxSteps, _ := strconv.Atoi(cmd.Flag("max_steps").Value.String())
	logInterval, _ := strconv.Atoi(cmd.Flag("log_interval").Value.String())
	saveInterval, _ := strconv.Atoi(cmd.Flag("save_interval").Value.String())
	return &TrainConfig{
		BatchSize:       batchSize,
		ImageChannels:   imageChannels,
		ImageHeight:     imageHeight,
		ImageWidth:      imageWidth,
		LatentChannels:  latentChannels,
		LatentHeight:    latentHeight,
		LatentWidth:     latentWidth,
		NumEpochs:       numEpochs,
		MaxSteps:        maxSteps,
		LogInterval:     logInterval,
		SaveInterval:    saveInterval,
		OutputDir:       cmd.Flag("output_dir").Value.String(),
		ModelPath:       cmd.Flag("model_path").Value.String(),
		TrainingDataDir: cmd.Flag("training_data_dir").Value.String(),
		CommandStepLogger: &CommandStepLogger{
			Cmd: cmd,
		},
	}
}

type Dataset interface {
	Batches() <-chan training.Batch
}

type Logger interface {
	LogStep(step int, epoch int, loss float32)
}

type CommandStepLogger struct {
	Cmd *cobra.Command
}

func (c *CommandStepLogger) LogStep(step int, epoch int, loss float32) {
	c.Cmd.Printf("Step %d epoch %d loss %f\n", step, epoch, loss)
}
