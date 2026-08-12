package components

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Paddle
}

func LoadPlayer(width, height, speed float32, color rl.Color) *Player {
	return &Player{
		Paddle: Paddle{
			Shape: rl.Rectangle{
				X: 75,
				Y: float32(rl.GetScreenHeight() / 2) - width / 2,
				Width: width,
				Height: height,
			},
			color: color,
			speed: speed,
			Score: 0,
		},
	}
}

func (p *Player) Move(dt float32) {
	if rl.IsKeyDown(rl.KeyW) {
		p.Shape.Y -= p.speed * dt
	}

	if rl.IsKeyDown(rl.KeyS) {
		p.Shape.Y += p.speed * dt
	}
}