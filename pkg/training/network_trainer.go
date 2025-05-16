// network_trainer.go
package training

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

// NetworkTrainer orchestrates the training loop
type NetworkTrainer struct {
	Model             *StableDiffusionModel
	Dataset           Dataset
	Config            *TrainConfig
	CommandStepLogger *CommandStepLogger
	Step              int
	Epoch             int
}

// NewNetworkTrainer creates a new instance of the trainer
func NewNetworkTrainer(cmd *cobra.Command, model *StableDiffusionModel, dataset Dataset, config *TrainConfig) *NetworkTrainer {
	return &NetworkTrainer{
		Model:             model,
		Dataset:           dataset,
		Config:            config,
		CommandStepLogger: &CommandStepLogger{Cmd: cmd},
	}
}

// Train kicks off the training loop
func (nt *NetworkTrainer) Train() error {
	startTime := time.Now()

	for epoch := 0; epoch < nt.Config.NumEpochs; epoch++ {
		nt.Epoch = epoch + 1
		nt.CommandStepLogger.Cmd.Printf("Starting epoch %d/%d", nt.Epoch, nt.Config.NumEpochs)

		for b := range nt.Dataset.Batches() {
			nt.Step++

			// Forward Pass
			latents := nt.Model.EncodeLatents(b.Images)
			pred := nt.Model.Forward(latents, b.Timestep, b.Conditioning)

			// Backward Pass
			grad := nt.Model.Backward(b.TargetNoise, pred)
			nt.Model.Step(grad)

			// Loss Calculation
			loss := computeMSE(pred, b.TargetNoise)

			if nt.Step%nt.Config.LogInterval == 0 {
				nt.CommandStepLogger.LogStep(nt.Step, nt.Epoch, loss)
			}

			// Checkpointing
			if nt.Config.SaveInterval > 0 && nt.Step%nt.Config.SaveInterval == 0 {
				filename := fmt.Sprintf("%s/model_step_%d.ckpt", nt.Config.OutputDir, nt.Step)
				_ = nt.Model.SaveCheckpoint(filename)
				nt.CommandStepLogger.Cmd.Printf("Saved checkpoint: %s", filename)
			}

			// Early exit
			if nt.Step >= nt.Config.MaxSteps {
				nt.CommandStepLogger.Cmd.Println("Reached max training steps.")
				break
			}
		}
		if nt.Step >= nt.Config.MaxSteps {
			break
		}
	}

	nt.CommandStepLogger.Cmd.Printf("Training finished in %s", time.Since(startTime))
	return nil
}

func computeMSE(a, b []float32) float32 {
	var sum float32
	for i := range a {
		diff := a[i] - b[i]
		sum += diff * diff
	}
	return sum / float32(len(a))
}
