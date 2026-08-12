package systems

import (
	"arcade-pong/components"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func UpdateScore(player *components.Player, opponent *components.Opponent, ball *components.Ball) {
	if ball.X < 0 {
		ball.X = float32(rl.GetScreenWidth() / 2) // IMPROVE BALL RESPAWN
		ball.Y = float32(rl.GetScreenHeight() / 2)
		opponent.AddScore()
	}
	
	if ball.X > float32(rl.GetScreenWidth()) {
		ball.X = float32(rl.GetScreenWidth() / 2)
		ball.Y = float32(rl.GetScreenHeight() / 2)
		player.AddScore()
	}
}

func EndGame(player, opponent components.Paddle) bool {
	if player.Score == 5 || opponent.Score == 5 {
		return true
	}

	return false
}