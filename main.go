package main

import (
	"image"
	"image/color"
	"log"
	"net/http"
	"os"

	"image/jpeg"
	"image/png"
	"math"
	"time"
)

func main() {
	args := os.Args

	// set timeout to 5 seconds.
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	//lookup image URL
	rawImage, err := client.Get(args[1])
	//e.g. https://2.bp.blogspot.com/-C3L7c1Y1ck4/UZOk1ldMiBI/AAAAAAAACs4/KtymYa4tM78/s1600/Screenshot+2013-05-15+at+12.38.58+AM.jpg
	if err != nil {
		log.Fatal(err)
	}
	defer rawImage.Body.Close()

	// Decode pallated and non-pallated images
	img, _, err := image.Decode(rawImage.Body)

	if err != nil {
		// handle error
		log.Fatal(err)
	}

	// Convert image to grayscale
	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			imageColor := img.At(x, y)

			red, green, blue, _ := imageColor.RGBA()

			calcRed := math.Pow(float64(red), 2.2)
			calcGreen := math.Pow(float64(green), 2.2)
			calcBlue := math.Pow(float64(blue), 2.2)

			//custom algorithm
			newColor := math.Pow(0.2125*calcRed+0.7154*calcGreen+0.0721*calcBlue, 1/2.2)

			intColor := uint16(newColor + 0.5)

			//divide by 8
			grayColor := color.Gray{uint8(intColor >> 8)}
			grayImg.Set(x, y, grayColor)

			//grayImg.Set(x, y, img.At(x, y))
		}
	}

	//detect the type of image
	switch img.(type) {

	case *image.YCbCr:
		file, err := os.Create("gray.jpg")
		if err != nil {
			// handle error
			log.Fatal(err)
		}
		defer file.Close()

		if err := jpeg.Encode(file, grayImg, nil); err != nil {
			log.Fatal(err)
		}

	case *image.RGBA:
		// Working with grayscale image, e.g. convert to png
		file, err := os.Create("gray.png")
		if err != nil {
			// handle error
			log.Fatal(err)
		}
		defer file.Close()

		if err := png.Encode(file, grayImg); err != nil {
			log.Fatal(err)
		}

	}
}
