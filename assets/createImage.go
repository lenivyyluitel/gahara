// this file is for creating a test image

package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
)

func main() {
	n := flag.Int("n", 0, "Number of images")
	flag.Parse()

	rect := image.Rect(0, 0, 600, 120)
	img := image.NewRGBA(rect)

	for z := 0; z < *n; z++ {
		for y := 0; y < 120; y++ {
			for x := 0; x < 600; x++ {
				img.Set(x, y, color.RGBA{
					R: uint8(x+y)>>z - 10,
					G: uint8(x+y)>>z - 10,
					B: uint8(x+y)>>z - 10,
					A: 255,
				})
			}
		}

		var file string = "image_" + strconv.Itoa(z) + ".png"
		f, err := os.Create(file)
		if err != nil {
			log.Panic(err)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {

			}
		}(f)
		err = png.Encode(f, img)
		if err != nil {
			log.Panic(err)
		}
	}
}
