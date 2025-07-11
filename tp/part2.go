package tp

import (
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"log"
	"math"
)

func convolution2D(img *image.NRGBA, kernel [][]float64) *image.Gray {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	output := image.NewGray(image.Rect(0, 0, width, height))

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			var temp float64

			for ky := -1; ky <= 1; ky++ {
				for kx := -1; kx <= 1; kx++ {
					px := img.NRGBAAt(x+kx, y+ky)
					gray := float64(px.R)
					temp += gray * kernel[ky+1][kx+1]
				}
			}

			// Clamping
			temp = math.Max(0, math.Min(255, temp))
			output.SetGray(x, y, color.Gray{Y: uint8(temp)})
		}
	}

	return output
}

func aplicaLimite(img *image.Gray) *image.Gray {
	bounds := img.Bounds()
	out := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			px := img.GrayAt(x, y).Y
			if px < 128 {
				out.SetGray(x, y, color.Gray{Y: 0})
			} else {
				out.SetGray(x, y, color.Gray{Y: 255})
			}
		}
	}
	return out
}
func Part2() {
	img, err := imaging.Open("fuji.jpg")
	if err != nil {
		log.Fatal(err)
	}
	grayNRGBA := imaging.Grayscale(img)
	imaging.Save(grayNRGBA, "fuji-cinza.jpg")
	// Kernel de relevo
	kRel := [][]float64{
		{-2, 1, 0},
		{-1, 1, 1},
		{0, 1, 2},
	}
	// Kernel de detecção de bordas
	kBordas := [][]float64{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}

	// Aplica filtro
	relevo := convolution2D(grayNRGBA, kRel)
	bordas := convolution2D(grayNRGBA, kBordas)

	bordas = aplicaLimite(bordas)

	// Salvar imagens
	imaging.Save(relevo, "fuji_relevo.jpg")
	imaging.Save(bordas, "fuji_bordas.jpg")
}
