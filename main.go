package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main()  {
	f, err := os.Open("assets/screenshot.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		panic(err)
	}

	err = invertImgColor(img)
	if err != nil {
		panic(err)
	}
}

func invertImgColor(img image.Image) error {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	rgbaScale := image.NewRGBA(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{w, h},
	})
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			imgColor := img.At(x, y)
			rr, gg, bb, aa := imgColor.RGBA()
			var max int64 = int64(int(255))
			nr := max - int64(int8(rr))
			ng := max - int64(int8(gg))
			nb := max - int64(int8(bb))
			invertColor := color.RGBA{
				R: uint8(nr),
				G: uint8(ng),
				B: uint8(nb),
				A: uint8(aa),
			}
			rgbaScale.Set(x, y, invertColor)
		}
	}

	newF, err := os.Create("assets/invertScreenshot.png")
	if err != nil {
		return err
	}
	defer newF.Close()
	png.Encode(newF, rgbaScale)
	return nil
}

func toGrayScaleImg(img image.Image) error {
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	grayScale := image.NewGray(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{w, h},
	})
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			imgColor := img.At(x, y)
			rr, gg, bb, _ := imgColor.RGBA()
			r := math.Pow(float64(rr), 2.2)
			g := math.Pow(float64(gg), 2.2)
			b := math.Pow(float64(bb), 2.2)
			m := math.Pow(0.2125*r+0.7154*g+0.0721*b, 1/2.2)
			Y := uint16(m + 0.5)
			grayColor := color.Gray{uint8(Y >> 8)}
			grayScale.Set(x, y, grayColor)
		}
	}

	newF, err := os.Create("assets/grayScreenshot.png")
	if err != nil {
		return err
	}
	defer newF.Close()
	png.Encode(newF, grayScale)
	return nil
}
