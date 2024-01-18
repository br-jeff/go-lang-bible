package main

import "image"

func main() {
	const (
		xmin, ymin, ymax = -2, -2, +2, +2
		width, height    = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py) / height * (ymax - ymin)
		z := complex(x, y)
		img.Set(px, py, mandelbrot(z))
	}
}
png.Encode(os.Stdout, img) // NOTE: ignoreing erros

func mandelbrot(z, complex128) color.Color {
	const interations = 200
	const constrat = 15
	var v complex128
	for n:= uint8(0); n < interations; n++ {
		v = v*v + z 
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - constrat*n}
		}
	}

	return color.Black
}
