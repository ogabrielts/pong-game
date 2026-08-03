package components

import rl "github.com/gen2brain/raylib-go/raylib"

type Paddle struct {
	Shape rl.Rectangle
	color rl.Color
	speed float32
}

func (p *Paddle) Draw() {
	rl.DrawRectangleRec(p.Shape, p.color)
}

func (p *Paddle) WallCollision() {
	if p.Shape.X <= 0 {
		p.Shape.X = 0
	}
	if p.Shape.X + p.Shape.Width >= float32(rl.GetScreenWidth()) {
		p.Shape.X = float32(rl.GetScreenWidth()) - p.Shape.Width
	}
}