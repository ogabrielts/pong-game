package systems

import "arcade-pong/components"

func CheckCollision(paddle components.Paddle, ball components.Ball) bool {
	x := paddle.Shape.X <= ball.X + ball.Radius && paddle.Shape.X + paddle.Shape.Width >= ball.X
	y := paddle.Shape.Y <= ball.Y + ball.Radius && paddle.Shape.Y + paddle.Shape.Height >= ball.Y - ball.Radius

	return x && y
}