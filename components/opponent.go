package components

import rl "github.com/gen2brain/raylib-go/raylib"

type Opponent struct {
	Paddle
}

func LoadOpponent(width, height, speed float32, color rl.Color) *Opponent {
	return &Opponent{
		Paddle: Paddle{
			shape: rl.Rectangle{
				X:      float32(rl.GetScreenWidth()/2) - width/2,
				Y:      float32(rl.GetScreenHeight() - 75) - height,
				Width:  width,
				Height: height,
			},
			color: color,
			speed: speed,
		},
	}
}

func (o *Opponent) Move(dt float32) {
	o.shape.X += o.speed * dt

	if o.shape.X <= 0 || o.shape.X + o.shape.Width >= float32(rl.GetScreenWidth()) {
		o.speed = -o.speed
	}
}