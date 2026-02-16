package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 800, "Atom Modeling")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	f := NewField()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		f.Draw()

		rl.EndDrawing()

		f.Next()
	}
}
