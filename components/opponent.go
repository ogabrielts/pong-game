package components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Opponent struct {
	Paddle
}

const delay float32 = 0.8
const errorMargin float32 = 20

func LoadOpponent(width, height, speed float32, color rl.Color) *Opponent {
	return &Opponent{
		Paddle: Paddle{
			Shape: rl.Rectangle{
				Y: float32(rl.GetScreenHeight() / 2) - height / 2,
				X: float32(rl.GetScreenWidth() - 75) - width,
				Width: width,
				Height: height,
			},
			color: color,
			speed: speed,
			Score: 0,
		},
	}
}

func (o *Opponent) Draw() {
	o.DrawPaddle()
	o.DrawScore(int32(rl.GetScreenWidth() / 2) + 20, 20)
}

func (o *Opponent) Move(ballY, ballX, dt float32) {
	o.WallCollision()

	paddleCenter := o.Shape.Y + (o.Shape.Height / 2)
	halfscreen := float32(rl.GetScreenWidth() / 2)
	middle := float32(rl.GetScreenHeight() / 2) - (o.Shape.Height / 2)

	if ballX < halfscreen {
		idle := rl.Vector2MoveTowards(rl.Vector2{X: o.Shape.X, Y: o.Shape.Y}, rl.Vector2{X: o.Shape.X, Y: middle}, (o.speed * dt) * delay)

		o.Shape.Y = idle.Y

	} else {
		if ballY > paddleCenter + errorMargin {
			o.Shape.Y += (o.speed * dt) * delay
		} else if ballY < paddleCenter - errorMargin {
			o.Shape.Y -= (o.speed * dt) * delay
		}	
	}
}