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

	player := components.LoadPlayer(100, 5, 250, rl.White)
	opponent := components.LoadOpponent(100, 5, 250, rl.White)
	ball := components.LoadBall(4, 400, rl.White)

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)

		ball.Draw()
		ball.Move(dt)
		ball.WallCollision()

		if ball.Y < 0 {
			ball.X = float32(rl.GetScreenWidth() / 2)
			ball.Y = float32(rl.GetScreenHeight() / 2)
			player.RemoveLive()
		}
		
		if ball.Y > float32(rl.GetScreenHeight()) {
			ball.X = float32(rl.GetScreenWidth() / 2)
			ball.Y = float32(rl.GetScreenHeight() / 2)
			opponent.RemoveLive()
		}

		if player.Lives <= 0 || opponent.Lives <= 0 {
			break
		}

		player.Draw()
		player.DrawScore(20, 20)
		player.Move(dt)
		player.WallCollision()

		opponent.Draw()
		opponent.DrawScore(20, int32(rl.GetScreenHeight()) - 50)
		opponent.Move(ball.X, dt)
		opponent.WallCollision()

		if systems.CheckCollision(player.Paddle, *ball) {
			hitOffset := ball.X - (player.Shape.X + player.Shape.Width / 2)
			normalized := hitOffset / (player.Shape.Width / 2)
			maxAngle := 0.75

			ball.Vel.X = normalized * float32(maxAngle) * ball.Speed
			ball.Vel.Y *= -1
		}

		if systems.CheckCollision(opponent.Paddle, *ball) {
			hitOffset := ball.X - (opponent.Shape.X + opponent.Shape.Width / 2)
			normalized := hitOffset / (opponent.Shape.Width / 2)
			maxAngle := 0.75

			ball.Vel.X = normalized * float32(maxAngle) * ball.Speed
			ball.Vel.Y *= -1
		}

		rl.EndDrawing()
	}
}