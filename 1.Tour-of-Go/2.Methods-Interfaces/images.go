package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// Rectángulo: de (0,0) a (100,100)
	rect := image.Rect(0, 0, 100, 100)

	// RGBA concreta (implementa image.Image y además tiene Set)
	img := image.NewRGBA(rect)

	// Pintar un gradiente simple
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			// r,g,b según posición; A=255 (opaco)
			c := color.RGBA{R: uint8(x * 255 / 99), G: uint8(y * 255 / 99), B: 160, A: 255}
			img.Set(x, y, c)
		}
	}

	// Guardar a disco (PNG)
	f, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}
