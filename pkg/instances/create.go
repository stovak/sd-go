package instances

import (
	"errors"
	"fmt"
	"os"
	"path"

	sd "github.com/seasonjs/stable-diffusion"
	"github.com/spf13/cobra"
)

func CreateModelInstance(cmd *cobra.Command) (*sd.Model, error) {
	cmd.Printf("Creating model instance")
	// Configure me...
	mp, err := cmd.Flags().GetString("model_path")
	if err != nil {
		return nil, errors.New(fmt.Sprint("unable to get model_path flag, err:", err))
	}
	if _, exist := os.Stat(mp); exist != nil {
		return nil, errors.New(fmt.Sprint("model path does not exist, err:", err))
	}
	cmd.Printf("Model Path: %s \n", mp)
	cp, err := cmd.Flags().GetString("checkpoint_filename")
	if err != nil {
		return nil, errors.New(fmt.Sprint("unable to get checkpoint_filename flag, err:", err))
	}
	if _, err = os.Stat(path.Join(mp, cp)); err != nil {
		return nil, errors.New(fmt.Sprintf("checkpoint file does not exist, err: %s", path.Join(mp, cp)))
	}

	options := sd.DefaultOptions
	model, err := sd.NewAutoModel(options)
	if err != nil {
		return nil, err
	}
	defer model.Close()
	cmd.Printf("options: %+v", options)

	err = model.LoadFromFile(path.Join(mp, cp))
	return model, err
}
