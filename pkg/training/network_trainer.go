// network_trainer.go
package training

import (
	"fmt"
	"github.com/stovak/sdgo/pkg/config"
	"log"
	"time"
)

type NetworkTrainer struct {
	Model   *StableDiffusionModel
	Dataset config.Dataset
	Config  *config.TrainConfig
	Logger  config.Logger
	Step    int
	Epoch   int
}

func NewNetworkTrainer(model *StableDiffusionModel, dataset config.Dataset, config *config.TrainConfig, logger config.Logger) *NetworkTrainer {
	return &NetworkTrainer{
		Model:   model,
		Dataset: dataset,
		Config:  config,
		Logger:  logger,
	}
}

func (nt *NetworkTrainer) Train() {
	startTime := time.Now()

	for epoch := 0; epoch < nt.Config.NumEpochs; epoch++ {
		nt.Epoch = epoch + 1
		log.Printf("Starting epoch %d/%d", nt.Epoch, nt.Config.NumEpochs)

		for batch := range nt.Dataset.Batches() {
			nt.Step++

			latents := make([]float32, nt.Config.BatchSize*nt.Config.LatentChannels*nt.Config.LatentHeight*nt.Config.LatentWidth)
			err := nt.Model.Encode(batch.Images, latents, nt.Config)
			if err != nil {
				log.Fatalf("Encode error: %v", err)
			}

			pred := make([]float32, len(latents))
			err = nt.Model.Forward(latents, batch.Timestep, batch.Conditioning, pred, nt.Config)
			if err != nil {
				log.Fatalf("Forward error: %v", err)
			}

			loss := computeMSE(pred, batch.TargetNoise)

			if nt.Step%nt.Config.LogInterval == 0 {
				nt.Logger.LogStep(nt.Step, nt.Epoch, loss)
			}

			if nt.Config.SaveInterval > 0 && nt.Step%nt.Config.SaveInterval == 0 {
				path := fmt.Sprintf("%s/model_step_%d.ckpt", nt.Config.OutputDir, nt.Step)
				nt.Model.SaveCheckpoint(path)
			}

			if nt.Step >= nt.Config.MaxSteps {
				break
			}
		}
		if nt.Step >= nt.Config.MaxSteps {
			break
		}
	}
	log.Printf("Training complete in %s", time.Since(startTime))
}

func computeMSE(a, b []float32) float32 {
	var sum float32
	for i := range a {
		diff := a[i] - b[i]
		sum += diff * diff
	}
	return sum / float32(len(a))
}
