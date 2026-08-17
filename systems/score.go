package systems

import (
	"arcade-pong/components"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func UpdateScore(player *components.Player, opponent *components.Opponent, ballX float32) bool {
	if ballX < 0 {
		opponent.AddScore()
		return true
	}
	
	if ballX > float32(rl.GetScreenWidth()) {
		player.AddScore()
		return true
	}

	return false
}

func EndGame(player, opponent components.Paddle) bool {
	if player.Score == 5 || opponent.Score == 5 {
		return true
	}

	return false
}