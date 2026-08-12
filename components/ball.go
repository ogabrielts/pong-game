package components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	X float32
	Y float32
	Radius float32
	Speed float32
	Vel rl.Vector2
	color rl.Color
}

func LoadBall(radius, speed float32, color rl.Color) *Ball {
	return &Ball{
		X: float32(rl.GetScreenWidth() / 2),
		Y: float32(rl.GetScreenHeight() / 2),
		Radius: radius,
		Speed: speed,
		Vel: rl.Vector2{
			X: speed,
			Y: 0,
		},
		color: color,
	}
}

func (b *Ball) Draw() {
	rl.DrawCircle(int32(b.X), int32(b.Y), b.Radius, b.color)
}

func (b *Ball) Move(dt float32) {
	b.Y += b.Vel.Y * dt
	b.X -= b.Vel.X * dt
}

func (b *Ball) WallCollision() {
	if (b.Y - b.Radius) <= 0 || (b.Y + b.Radius) >= float32(rl.GetScreenHeight()) {
		b.Vel.Y *= -1
	}
}