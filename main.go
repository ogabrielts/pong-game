package main

import (
	"arcade-pong/components"
	"arcade-pong/systems"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1024, 720, "Arcade Pong")
	defer rl.CloseWindow()

	rl.SetTargetFPS(240)

	// Load player, opponent, and ball
	player := components.LoadPlayer(5, 100, 250, rl.White)
	opponent := components.LoadOpponent(5, 100, 250, rl.White)
	ball := components.LoadBall(4, 400, rl.White)

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()

		rl.BeginDrawing()

		// Background
		rl.ClearBackground(rl.DarkGray)
		rl.DrawRectangle(int32((rl.GetScreenWidth()) / 2) - 1, 0, 2, int32(rl.GetScreenHeight()), rl.Gray)

		// Check Score
		systems.UpdateScore(player, opponent, ball)

		// Ball logic
		ball.Draw()
		ball.Move(dt)
		ball.WallCollision()

		// Player logic
		player.Draw()
		player.DrawScore(int32(rl.GetScreenWidth() / 2) - 30, 20)
		player.Move(dt)
		player.WallCollision()

		// Check ball-paddle collision for player
		if systems.CheckCollision(player.Paddle, *ball) {
			hitOffset := ball.Y - (player.Shape.Y + player.Shape.Height / 2)
			normalized := hitOffset / (player.Shape.Height / 2)
			maxAngle := 0.75

			ball.Vel.Y = normalized * float32(maxAngle) * ball.Speed
			ball.Vel.X *= -1
		}

		// Opponent logic
		opponent.Draw()
		opponent.DrawScore(int32(rl.GetScreenWidth() / 2) + 20, 20)
		opponent.Move(ball.Y, dt)
		opponent.WallCollision()

		// Check ball-paddle collision for opponent
		if systems.CheckCollision(opponent.Paddle, *ball) {
			hitOffset := ball.Y - (opponent.Shape.Y + opponent.Shape.Height / 2)
			normalized := hitOffset / (opponent.Shape.Height / 2)
			maxAngle := 0.75

			ball.Vel.Y = normalized * float32(maxAngle) * ball.Speed
			ball.Vel.X *= -1
		}


		// End loop if score == 5
		if systems.EndGame(player.Paddle, opponent.Paddle) {
			break
		}

		rl.EndDrawing()
	}
}