package field

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (f *Field) Draw() {
	imageData := make([]byte, L*L)
	for i := range f.values {
		for j := range f.values[i] {
			cF := squareOfAbs(f.values[i][j])
			c := uint8(min(cF/Dx/Dx * 64, 255))

			imageData[i*L+j] = c
		}
	}
	img := rl.NewImage(imageData, L, L, 1, rl.UncompressedGrayscale)
	texture := rl.LoadTextureFromImage(img)
	rl.DrawTextureEx(texture, rl.NewVector2(0, 0), 0, 800 / L, rl.White)
}

func (f *Field) DrawRB() {
	imageData := make([]byte, 3*L*L)
	for i := range f.values {
		for j := range f.values[i] {
			rF := (real(f.values[i][j]))
			r := uint8(max(min(128 + rF*128/Dx, 255), 0))

			bF := (imag(f.values[i][j]))
			b := uint8(max(min(128 + bF*128/Dx, 255), 0))

			imageData[3*(i*L+j)] = r
			imageData[3*(i*L+j) + 1] = 128
			imageData[3*(i*L+j)+2] = b
		}
	}
	img := rl.NewImage(imageData, L, L, 1, rl.UncompressedR8g8b8)
	texture := rl.LoadTextureFromImage(img)
	rl.DrawTextureEx(texture, rl.NewVector2(0, 0), 0, 800 / L, rl.White)
}
