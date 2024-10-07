package utils

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadImage(path string) *ebiten.Image {	
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}