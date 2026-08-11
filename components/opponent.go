package components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Opponent struct {
	Paddle
}

const delay float32 = 0.85

func LoadOpponent(width, height, speed float32, color rl.Color) *Opponent {
	return &Opponent{
		Paddle: Paddle{
			Shape: rl.Rectangle{
				X:      float32(rl.GetScreenWidth()/2) - width/2,
				Y:      float32(rl.GetScreenHeight() - 75) - height,
				Width:  width,
				Height: height,
			},
			color: color,
			speed: speed,
			Lives: 3,
		},
	}
}

func (o *Opponent) Move(ballPos, dt float32) {
	paddleCenter := o.Shape.X + (o.Shape.Width / 2)

	if ballPos > paddleCenter {
		o.Shape.X += o.speed * (dt * delay)
	} else if ballPos < paddleCenter {
		o.Shape.X -= o.speed * (dt * delay)
	}
}