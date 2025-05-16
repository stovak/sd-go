package training

import (
	"image"
	"os"
	"path/filepath"
	"slices"
)

var DefaultImageExtensions = []string{
	".jpg",
	".jpeg",
	".png",
	".gif",
	".webp",
}

type Batch struct {
	Images       []image.Image
	Timestep     int
	Conditioning []float32
	TargetNoise  []float32
}

func (b *Batch) AddImage(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}
	b.Images = append(b.Images, img)
	return nil
}

func LoadBatchFromDirectory(directory string, fileExtensions []string) (Batch, error) {
	var b Batch
	files, err := os.ReadDir(directory)
	if err != nil {
		return b, err
	}
	for _, f := range files {
		if f.IsDir() {
			// Load the files from inside the directory
			b2, err := LoadBatchFromDirectory(filepath.Join(directory, f.Name()), fileExtensions)
			if err != nil {
				return Batch{}, err
			}
			b.Images = append(b.Images, b2.Images...)
		}
		// search the extensions slice for the current extension
		if slices.Contains(fileExtensions, filepath.Ext(f.Name())) {
			err = b.AddImage(filepath.Join(directory, f.Name()))
			if err != nil {
				return Batch{}, err
			}
		}

	}
	return b, nil
}
