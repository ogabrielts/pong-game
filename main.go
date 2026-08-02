package main

import (
	"arcade-pong/components"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(720, 840, "Arcade Pong")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	player := components.LoadPlayer(150, 10, 150, rl.White)
	opponent := components.LoadOpponent(150, 10, 150, rl.White)
	ball := components.LoadBall(8, 150, rl.White)

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)

		ball.Draw()

		player.Draw()
		player.Move(dt)
		player.WallCollision()

		opponent.Draw()
		opponent.Move(dt)
		opponent.WallCollision()

		rl.EndDrawing()
	}
}