package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// Define the Image type with width and height.
type Image struct {
	width, height int
}

// Bounds returns the rectangle with the image bounds.
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

// ColorModel returns the color model of the image, which is RGBA.
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// At returns the color of the pixel at (x, y).
func (img Image) At(x, y int) color.Color {
	v := uint8((x ^ y) % 256) // Adjust this function to experiment with patterns
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{width: 256, height: 256}
	//pic.ShowImage(m)
	// Create a file to save the image
	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Encode and save the image to the file
	err = png.Encode(file, m)
	if err != nil {
		panic(err)
	}

	println("Image saved as output.png")
}
