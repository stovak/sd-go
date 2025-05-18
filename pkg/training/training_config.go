package training

import (
	"github.com/spf13/cobra"
	"strconv"
)

type TrainConfig struct {
	// Now she's back in my atmosphere
	// with drops of jupiter in her hair, heh-heh

	// === Model Input Shapes ===
	BatchSize     int // e.g., 1, 2, 4
	ImageChannels int // typically 3 for RGB
	ImageHeight   int // e.g., 512
	ImageWidth    int // e.g., 512

	LatentChannels int // typically 4 (for SD 1.5)
	LatentHeight   int // usually ImageHeight / 8 (e.g., 64)
	LatentWidth    int // usually ImageWidth / 8 (e.g., 64)

	// === Training Parameters ===
	NumEpochs    int // e.g., 10
	MaxSteps     int // total training steps (overrides epochs if set)
	LogInterval  int // log every N steps
	SaveInterval int // save every N steps

	// === File System / Paths ===
	OutputDir string // where to write logs/models

	ModelPath string // path to .ckpt or .safetensors if used

	TrainingDataDir string

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
	Batches() <-chan Batch
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
