package instances

import (
	"errors"
	"fmt"
	"github.com/seasonjs/hf-hub/api"
	sd "github.com/seasonjs/stable-diffusion"
	"github.com/spf13/cobra"
)

func CreateModelInstance(cmd *cobra.Command) (*sd.Model, error) {
	// Configure me...
	mp, err := cmd.Flags().GetString("model_path")
	if err != nil {
		return nil, errors.New(fmt.Sprint("unable to get model_path flag, err:", err))
	}
	cp, err := cmd.Flags().GetString("checkpoint_filename")
	if err != nil {
		return nil, errors.New(fmt.Sprint("unable to get checkpoint_filename flag, err:", err))
	}

	options := sd.DefaultOptions
	// TODO: merge options from config file with defaults
	model, err := sd.NewAutoModel(options)
	if err != nil {
		return nil, errors.New(fmt.Sprint("unable to create auto model, err:", err))
	}
	defer model.Close()

	hapi, err := api.NewApi()
	if err != nil {
		return nil, errors.New(fmt.Sprint("unable to create api instance, err:", err))
	}
	modelPath, err := hapi.Model(mp).Get(cp)
	if err != nil {
		return nil, errors.New(fmt.Sprint("unable to get model path, err:", err))
	}

	err = model.LoadFromFile(modelPath)
	if err != nil {
		return nil, errors.New(fmt.Sprint("unable to load model file, err:", err))
	}
	return model, nil
}
