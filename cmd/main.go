package main

import (
	"atom-modeling/internal/field"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 800, "Atom Modeling")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	f := field.New(field.Config{
		Mass:                   50,
		ReducedPlancksConstant: 1.66,
	})

	drawRb := false

	for !rl.WindowShouldClose() {
		for {
			key := rl.GetKeyPressed()
			if key == 0 {
				break
			}
			if key == rl.KeyM {
				drawRb = !drawRb
			}

		}
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		if drawRb {
			f.DrawRB()
		} else {
			f.Draw()
		}

		frameTimeStr := fmt.Sprintf("Frame time: %.3f", rl.GetFrameTime())
		rl.DrawText(frameTimeStr, 0, 0, 20, rl.White)

		rl.EndDrawing()

		for range 4 {
			f.Next()
		}
	}
}
