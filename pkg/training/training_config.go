package training

import (
	"github.com/spf13/cobra"
)

type TrainConfig struct {
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

	// (Optional)
	ModelPath string // path to .ckpt or .safetensors if used

	CommandStepLogger *CommandStepLogger
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
