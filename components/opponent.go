package components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Opponent struct {
	Paddle
}

const delay float32 = 0.80

func LoadOpponent(width, height, speed float32, color rl.Color) *Opponent {
	return &Opponent{
		Paddle: Paddle{
			Shape: rl.Rectangle{
				Y:      float32(rl.GetScreenHeight()/2) - height/2,
				X:      float32(rl.GetScreenWidth() - 75) - width,
				Width:  width,
				Height: height,
			},
			color: color,
			speed: speed,
			Score: 0,
		},
	}
}

func (o *Opponent) Move(ballPos, dt float32) { // FIX AI PADDLE MOVEMENT
	paddleCenter := o.Shape.Y + (o.Shape.Height / 2)

	if ballPos > paddleCenter {
		o.Shape.Y += o.speed * (dt) * delay
	} else if ballPos < paddleCenter {
		o.Shape.Y -= o.speed * (dt) * delay
	}
}