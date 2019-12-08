package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
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

	// get layer with lowest numbers of zeros
	layerIndex := 0
	lowestZeroes := math.MaxInt64
	for layer := 0; layer < img.LayerCount(); layer++ {
		zeroDigits := countOccurancesInLayer(img, layer, '0')
		if zeroDigits < lowestZeroes {
			layerIndex = layer
			lowestZeroes = zeroDigits
		}
	}


	numberOfOnes := countOccurancesInLayer(img, layerIndex, '1')
	numberOfTwos := countOccurancesInLayer(img, layerIndex, '2')

	fmt.Println(numberOfOnes * numberOfTwos)
}

func countOccurancesInLayer(img *Image, layer int, needle byte) int {
	count := 0
	for x := 0; x < img.Width; x++ {
		for y := 0; y < img.Height; y++ {
			if img.GetPixel(x, y, layer) == needle {
				count++
			}
		}
	}
	return count
}