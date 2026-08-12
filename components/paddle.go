package components

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Paddle struct {
	Shape rl.Rectangle
	color rl.Color
	speed float32
	Score int
}

func (p *Paddle) Draw() {
	rl.DrawRectangleRec(p.Shape, p.color)
}

func (p *Paddle) DrawScore(x, y int32) {
	rl.DrawText(strconv.Itoa(p.Score), x, y, 20, rl.White)
}

func (p *Paddle) WallCollision() {
	if p.Shape.Y <= 0 {
		p.Shape.Y = 0
	}
	if p.Shape.Y + p.Shape.Height >= float32(rl.GetScreenHeight()) {
		p.Shape.Y = float32(rl.GetScreenHeight()) - p.Shape.Height
	}
}

func (p *Paddle) AddScore() {
	p.Score++
}