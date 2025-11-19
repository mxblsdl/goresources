package extras

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func CreateIcon() {
	// Create a 512x512 image for the icon
	size := 512
	img := image.NewRGBA(image.Rect(0, 0, size, size))

	// Colors
	blue := color.RGBA{41, 128, 185, 255}
	white := color.RGBA{255, 255, 255, 255}
	darkBlue := color.RGBA{25, 80, 120, 255}

	// Draw background
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			img.Set(x, y, blue)
		}
	}

	// Draw border
	borderWidth := 10
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if x < borderWidth || x >= size-borderWidth ||
				y < borderWidth || y >= size-borderWidth {
				img.Set(x, y, white)
			}
		}
	}

	// Draw a simple "M" for Monitor in the center
	centerX := size / 2
	centerY := size / 2
	letterSize := 200
	thickness := 30

	// Left vertical line of M
	for y := centerY - letterSize/2; y < centerY+letterSize/2; y++ {
		for x := centerX - letterSize/2; x < centerX-letterSize/2+thickness; x++ {
			if x >= 0 && x < size && y >= 0 && y < size {
				img.Set(x, y, darkBlue)
			}
		}
	}

	// Right vertical line of M
	for y := centerY - letterSize/2; y < centerY+letterSize/2; y++ {
		for x := centerX + letterSize/2 - thickness; x < centerX+letterSize/2; x++ {
			if x >= 0 && x < size && y >= 0 && y < size {
				img.Set(x, y, darkBlue)
			}
		}
	}

	// Left diagonal of M
	for i := 0; i < letterSize/2; i++ {
		for t := 0; t < thickness; t++ {
			x := centerX - letterSize/2 + i + t
			y := centerY - letterSize/2 + i
			if x >= 0 && x < size && y >= 0 && y < size {
				img.Set(x, y, darkBlue)
			}
		}
	}

	// Right diagonal of M
	for i := 0; i < letterSize/2; i++ {
		for t := 0; t < thickness; t++ {
			x := centerX + letterSize/2 - i - t
			y := centerY - letterSize/2 + i
			if x >= 0 && x < size && y >= 0 && y < size {
				img.Set(x, y, darkBlue)
			}
		}
	}

	// Save to file
	f, err := os.Create("icon.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}

	println("Icon created: icon.png")
}
