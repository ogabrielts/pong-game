package main

import (
	"arcade-pong/components"
	"arcade-pong/systems"

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
		ball.Move(dt)
		ball.WallCollision()

		player.Draw()
		player.Move(dt)
		player.WallCollision()

		opponent.Draw()
		opponent.Move(ball.X, dt)
		opponent.WallCollision()

		if systems.CheckCollision(player.Paddle, *ball) {
			hitOffset := ball.X - (player.Shape.X + player.Shape.Width / 2)
			normalized := hitOffset / (player.Shape.Width / 2)
			maxAngle := 0.75

			ball.Speed.X = normalized * float32(maxAngle) * ball.BaseSpeed
			ball.Speed.Y *= -1
		}

		if systems.CheckCollision(opponent.Paddle, *ball) {
			hitOffset := ball.X - (opponent.Shape.X + opponent.Shape.Width / 2)
			normalized := hitOffset / (opponent.Shape.Width / 2)
			maxAngle := 0.75

			ball.Speed.X = normalized * float32(maxAngle) * ball.BaseSpeed
			ball.Speed.Y *= -1
		}

		rl.EndDrawing()
	}
}