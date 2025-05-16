package training

/*
   #cgo LDFLAGS: -L./lib -lsd-abi
   #include <stdlib.h>
   #include "stable_diffusion_api.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// StableDiffusionModel is a wrapper around the C++ libsd-abi.dylib model
type StableDiffusionModel struct {
	ctx C.sd_context_t
}

// LoadModel loads the model from a given path
func LoadModel(modelPath string) (*StableDiffusionModel, error) {
	cPath := C.CString(modelPath)
	defer C.free(unsafe.Pointer(cPath))

	ctx := C.sd_load_model(cPath)
	if ctx == nil {
		return nil, fmt.Errorf("failed to load model at %s", modelPath)
	}
	return &StableDiffusionModel{ctx: ctx}, nil
}

// Free releases the C++ model
func (m *StableDiffusionModel) Free() {
	if m.ctx != nil {
		C.sd_free_model(m.ctx)
		m.ctx = nil
	}
}

// Encode encodes a batch of images to latents
func (m *StableDiffusionModel) Encode(images, latentsOut []float32, cfg *TrainConfig) error {
	if len(images) != cfg.BatchSize*cfg.ImageChannels*cfg.ImageHeight*cfg.ImageWidth {
		return fmt.Errorf("invalid image input shape")
	}
	if len(latentsOut) != cfg.BatchSize*cfg.LatentChannels*cfg.LatentHeight*cfg.LatentWidth {
		return fmt.Errorf("invalid latent output shape")
	}

	errCode := C.sd_encode_latents(
		m.ctx,
		(*C.float)(&images[0]),
		C.size_t(cfg.BatchSize),
		C.size_t(cfg.ImageChannels),
		C.size_t(cfg.ImageHeight),
		C.size_t(cfg.ImageWidth),
		(*C.float)(&latentsOut[0]),
	)
	if errCode != 0 {
		return fmt.Errorf("sd_encode_latents failed with code %d", errCode)
	}
	return nil
}

// Forward runs the U-Net forward pass
func (m *StableDiffusionModel) Forward(latents []float32, timestep int, conditioning []float32, output []float32, cfg *TrainConfig) error {
	if len(latents) != cfg.BatchSize*cfg.LatentChannels*cfg.LatentHeight*cfg.LatentWidth {
		return fmt.Errorf("invalid latent input shape")
	}
	if len(output) != len(latents) {
		return fmt.Errorf("output shape must match input latents shape")
	}

	errCode := C.sd_forward(
		m.ctx,
		(*C.float)(&latents[0]),
		C.size_t(cfg.BatchSize),
		C.size_t(cfg.LatentChannels),
		C.size_t(cfg.LatentHeight),
		C.size_t(cfg.LatentWidth),
		C.int(timestep),
		(*C.float)(&conditioning[0]),
		C.size_t(len(conditioning)),
		(*C.float)(&output[0]),
	)
	if errCode != 0 {
		return fmt.Errorf("sd_forward failed with code %d", errCode)
	}
	return nil
}

// SaveCheckpoint is a stub — if your C++ lib supports saving, call that here
func (m *StableDiffusionModel) SaveCheckpoint(path string) {
	// no-op for now unless `sd_save_checkpoint` exists
}
