package components

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	X float32
	Y float32
	Radius float32
	Speed float32
	Vel rl.Vector2
	color rl.Color
	SpawnTimer float32
}

func LoadBall(radius, speed float32, color rl.Color) *Ball {
	return &Ball{
		X: float32(rl.GetScreenWidth() / 2),
		Y: float32(rl.GetScreenHeight() / 2),
		Radius: radius,
		Speed: speed,
		Vel: rl.Vector2{
			X: randomizeX(speed),
			Y: randomizeY(speed),
		},
		color: color,
		SpawnTimer: 2.0,
	}
}

func randomizeX(speed float32) float32 {
	num := rand.Intn(2)

	if num == 0 {
		return speed
	} else {
		return -speed
	}
}

func randomizeY(speed float32) float32 {
	num := rand.Float32()
	return speed * (num - 0.5)
}

func (b *Ball) Draw() {
	rl.DrawCircle(int32(b.X), int32(b.Y), b.Radius, b.color)
}

func (b *Ball) Move(dt float32) {
	if b.SpawnTimer > 0 {
		b.SpawnTimer -= dt
	} else {
		b.Y += b.Vel.Y * dt
		b.X -= b.Vel.X * dt
	}

	b.wallCollision()
}

func (b *Ball) wallCollision() {
	if (b.Y - b.Radius) <= 0 || (b.Y + b.Radius) >= float32(rl.GetScreenHeight()) {
		b.Vel.Y *= -1
	}
}