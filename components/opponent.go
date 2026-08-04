package components

import rl "github.com/gen2brain/raylib-go/raylib"

type Opponent struct {
	Paddle
}

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
		},
	}
}

func (o *Opponent) Move(ballPos, dt float32) {
	if ballPos >= o.Shape.X + (o.Shape.Width / 2) {
		o.Shape.X += o.speed * dt
	}

	if ballPos <= o.Shape.X + (o.Shape.Width / 2) {
		o.Shape.X -= o.speed * dt
	}


}