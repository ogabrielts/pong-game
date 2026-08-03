package components

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Paddle
}

func LoadPlayer(width, height, speed float32, color rl.Color) *Player {
	return &Player{
		Paddle: Paddle{
			Shape: rl.Rectangle{
				X: float32(rl.GetScreenWidth() / 2) - width / 2,
				Y: 75,
				Width: width,
				Height: height,
			},
			color: color,
			speed: speed,
		},
	}
}

func (p *Player) Move(dt float32) {
	if rl.IsKeyDown(rl.KeyA) {
		p.Shape.X -= p.speed * dt
	}

	if rl.IsKeyDown(rl.KeyD) {
		p.Shape.X += p.speed * dt
	}
}