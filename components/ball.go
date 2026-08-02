package components

import rl "github.com/gen2brain/raylib-go/raylib"

type Ball struct {
	x int32
	y int32
	radius float32
	speed float32
	color rl.Color
}

func LoadBall(radius, speed float32, color rl.Color) *Ball {
	return &Ball{
		x: int32(rl.GetScreenWidth() / 2),
		y: int32(rl.GetScreenHeight() / 2),
		radius: radius,
		speed: speed,
		color: color,

	}
}

func (b *Ball) Draw() {
	rl.DrawCircle(b.x, b.y, b.radius, b.color)
}

