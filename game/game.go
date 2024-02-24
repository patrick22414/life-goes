package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/gen2brain/raylib-go/raygui"
	"github.com/patrick22414/life-goes/matrix"
)

const WW, WH = 800, 480
const NX, NY = 8, 8
const FPS = 60

func Start() {
	rl.InitWindow(WW, WH, "raylib - Conway's Game of Life")
	defer rl.CloseWindow()

	rl.SetTargetFPS(FPS)

	grid := Grid{
		nx:   NX,
		ny:   NY,
		size: 40,
		line: 4,
	}
	grid.CenterInScreen(WW, WH)

	board := matrix.NewRand(NX, NY)

	live := false
	buttonText := "Start"

	i := 0
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		grid.DrawGrid(rl.DarkGray)
		for x := range NX {
			for y := range NY {
				v := board.I(x, y)
				if v != 0 {
					grid.DrawCell(x, y, rl.Beige)
				}
			}
		}

		// GUI
		if raygui.Button(rl.Rectangle{X: 4, Y: 4, Width: 100, Height: 25}, buttonText) {
			live = !live
			if live {
				buttonText = "Pause"
			} else {
				buttonText = "Continue"
			}
		}
		if raygui.Button(rl.Rectangle{X: 124, Y: 4, Width: 100, Height: 25}, "Clear") {
			board = matrix.New(NX, NY)
		}
		if raygui.Button(rl.Rectangle{X: 244, Y: 4, Width: 100, Height: 25}, "Random") {
			board = matrix.NewRand(NX, NY)
		}

		if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
			mousePos := rl.GetMousePosition()
			x, y, in := grid.ScreenToGrid(mousePos)
			if in {
				board.Flip(x, y)
			}
		}

		if live {
			raygui.Label(rl.Rectangle{X: 4, Y: 4 + 25 + 4, Width: 100, Height: 25}, "live")
		}

		if i = (i + 1) % (FPS / 5); i == 0 && live {
			board = board.Next()
		}

		rl.EndDrawing()
	}
}
