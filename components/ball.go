package components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	X float32
	Y float32
	Radius float32
	BaseSpeed float32
	Speed rl.Vector2
	color rl.Color
}

func LoadBall(radius, speed float32, color rl.Color) *Ball {
	return &Ball{
		X: float32(rl.GetScreenWidth() / 2),
		Y: float32(rl.GetScreenHeight() / 2),
		Radius: radius,
		BaseSpeed: speed,
		Speed: rl.Vector2{
			X: 0,
			Y: speed,
		},
		color: color,

	}
}

func (b *Ball) Draw() {
	rl.DrawCircle(int32(b.X), int32(b.Y), b.Radius, b.color)
}

func (b *Ball) Move(dt float32) {
	b.Y -= b.Speed.Y * dt
	b.X += b.Speed.X * dt
}

func (b *Ball) WallCollision() {
	if (b.X - b.Radius) <= 0 || (b.X + b.Radius) >= float32(rl.GetScreenWidth()) {
		b.Speed.X *= -1
	}
}