package components

import rl "github.com/gen2brain/raylib-go/raylib"

type Paddle struct {
	shape rl.Rectangle
	color rl.Color
	speed float32
}

func (p *Paddle) Draw() {
	rl.DrawRectangleRec(p.shape, p.color)
}

func (p *Paddle) WallCollision() {
	if p.shape.X <= 0 {
		p.shape.X = 0
	}
	if p.shape.X + p.shape.Width >= float32(rl.GetScreenWidth()) {
		p.shape.X = float32(rl.GetScreenWidth()) - p.shape.Width
	}
}