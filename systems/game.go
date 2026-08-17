package systems

import (
	"arcade-pong/components"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	player *components.Player
	opponent *components.Opponent
	ball *components.Ball
}

func NewGame() *Game{
	return &Game{
		player: components.LoadPlayer(5, 100, 250, rl.White),
		opponent: components.LoadOpponent(5, 100, 250, rl.White),
		ball: components.LoadBall(4, 400, rl.White),
	}
}

func (game *Game) Update(dt float32) {
	// Check Score, remove the old ball, and spawn a new one at the same position
	if UpdateScore(game.player, game.opponent, game.ball.X) {
		game.ball = nil
		newBall := components.LoadBall(4, 400, rl.White)
		game.ball = newBall
	}

	// Move ball, player, and opponent
	game.ball.Move(dt)
	game.player.Move(dt)
	game.opponent.Move(game.ball.Y, game.ball.X, dt)

	// Ball-player collision
	if CheckCollision(game.player.Paddle, *game.ball) {
		hitOffset := game.ball.Y - (game.player.Shape.Y + game.player.Shape.Height / 2)
		normalized := hitOffset / (game.player.Shape.Height / 2)
		maxAngle := 0.75

		game.ball.Vel.Y = normalized * float32(maxAngle) * game.ball.Speed
		game.ball.Vel.X *= -1
	}

	// Ball-opponent collision
	if CheckCollision(game.opponent.Paddle, *game.ball) {
		hitOffset := game.ball.Y - (game.opponent.Shape.Y + game.opponent.Shape.Height / 2)
		normalized := hitOffset / (game.opponent.Shape.Height / 2)
		maxAngle := 0.75

		game.ball.Vel.Y = normalized * float32(maxAngle) * game.ball.Speed
		game.ball.Vel.X *= -1
	}

}

func (game *Game) Draw() {
	// Background
	rl.ClearBackground(rl.DarkGray)
	rl.DrawRectangle(int32((rl.GetScreenWidth()) / 2) - 1, 0, 2, int32(rl.GetScreenHeight()), rl.Gray)

	// Draw ball, player, opponent
	game.ball.Draw()
	game.player.Draw()
	game.opponent.Draw()
}

func (game *Game) End() bool {
	return EndGame(game.player.Paddle, game.opponent.Paddle)
}