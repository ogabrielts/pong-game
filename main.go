package main

import (
	"arcade-pong/systems"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type State int

const (
	Menu State = iota
	Play
	Replay
)

func main() {
	rl.InitWindow(1024, 720, "Arcade Pong")
	defer rl.CloseWindow()

	rl.SetTargetFPS(240)

	game := systems.NewGame()
	var state State = 0


	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()


		switch state {
		case Menu:
			if rl.IsKeyDown(rl.KeySpace) {
				state = 1
			}

		case Play:
			game.Update(dt)

		case Replay:
			if rl.IsKeyDown(rl.KeySpace) {
				state = 1
			}
		}

		rl.BeginDrawing()

		switch state {
		case Menu:
			rl.ClearBackground(rl.DarkBlue)
			rl.DrawText("Just Pong", 30, 100, 50, rl.White)
			rl.DrawText("Press [SPACE] to start", 30, 200, 30, rl.White)

		case Play:
			game.Draw()
			if game.End() {
				state = 2
			}

		case Replay:
			rl.ClearBackground(rl.DarkBlue)
			rl.DrawText("Game Over", 30, 100, 50, rl.White)
			rl.DrawText("Press [SPACE] to play again", 30, 200, 30, rl.White)

			game = systems.NewGame()
		}

		rl.EndDrawing()
	}
}