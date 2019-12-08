package main

import (
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

type Image struct {
	Width  int
	Height int
	Data []byte
}

func (i *Image) LayerCount() int {
	return len(i.Data) / (i.Width * i.Height)
}

func (i *Image) GetPixel(x, y, layer int) byte {
	idx := layer * (i.Width * i.Height) + y * i.Width + x
	return i.Data[idx]
}

func (i *Image) ToRGB() *image.RGBA {
	img := image.NewRGBA(image.Rectangle{Max: image.Point{X: i.Width, Y: i.Height}})

	for j := i.LayerCount() - 1; j >= 0; j-- {
		for x := 0; x < i.Width; x++ {
			for y := 0; y < i.Height; y++ {
				switch i.GetPixel(x, y, j) {
				case '0':
					img.Set(x, y, color.Black)
				case '1':
					img.Set(x, y, color.White)
				}
			}
		}
	}

	return img
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to load input: %v\n", err)
	}

	img := &Image{
		Width:  25,
		Height: 6,
		Data: b,
	}

	out := img.ToRGB()
	f, _ := os.Create("out.png")
	if err := png.Encode(f, out); err != nil {
		panic(err)
	}
}
