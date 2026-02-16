package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Field struct {
	values [][]complex128
}

const Dx = float64(0.01)
const Dt = float64(0.001)

const Dxc = complex128(Dx)
const Dtc = complex128(Dt)

func initValues() [][]complex128 {
	values := make([][]complex128, 800)
	for i := range values {
		values[i] = make([]complex128, 800)
	}
	return values
}

func NewField() Field {
	values := initValues()
	for i := range 800 {
		for j := range 800 {
			dist := (i-400)*(i-400) + (j-400)*(j-400)
			values[i][j] = complex(math.Exp(-float64(dist)/10000), 0)
				// cmplx.Exp(complex(0, float64(i)*0.001))
		}
	}
	f := Field{values}
	f.normalize()
	return f
}

func (f *Field) get(i, j int) complex128 {
	if i < 0 || j < 0 {
		return 0
	}
	if i >= 800 || j >= 800 {
		return 0
	}
	return f.values[i][j]
}

func (f *Field) Next() {
	next := initValues()

	a := complex(0, Dt/Dx/Dx*0.01)
	b := complex(0, Dt*10)

	for i := range f.values {
		for j := range f.values[i] {
			v := a*(f.get(i+1, j)+f.get(i, j+1)+f.get(i-1, j)+f.get(i, j-1)-4*f.get(i, j)) -
				b*f.get(i, j) +
				f.get(i, j)

			next[i][j] = v
		}
	}

	f.values = next
	f.normalize()
}

func (f *Field) Draw() {
	imageData := make([]byte, 800*800)
	for i := range f.values {
		for j := range f.values[i] {
			cF := squareOfAbs(f.values[i][j])
			c := uint8(min(cF*100, 255))
			imageData[i*800+j] = c
		}
	}
	img := rl.NewImage(imageData, 800, 800, 1, rl.UncompressedGrayscale)
	texture := rl.LoadTextureFromImage(img)
	rl.DrawTexture(texture, 0, 0, rl.White)
}

func squareOfAbs(c complex128) float64 {
	x := real(c)
	y := imag(c)
	return x*x + y*y
}

func (f *Field) normalize() {
	sum := float64(0)
	for i := range f.values {
		for j := range f.values[i] {
			sum += squareOfAbs(f.values[i][j]) * Dx * Dx
		}
	}

	if sum > 1 {
		fmt.Println(sum)
	}
	m := 1 / math.Sqrt(sum)

	for i := range f.values {
		for j := range f.values[i] {
			f.values[i][j] *= complex(m, 0)
		}
	}
}
