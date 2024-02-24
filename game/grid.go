package game

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Grid struct {
	nx, ny int32   // number of tiles
	size   int32   // tile size, including line
	line   int32   // line thickness
	xs, ys []int32 // line positions
}

func (g *Grid) WH() (w, h int32) {
	return g.nx*g.size + g.line, g.ny*g.size + g.line
}

func (g *Grid) CenterInScreen(screenW int32, screenH int32) {
	w, h := g.WH()

	x := (screenW - w - g.line) / 2
	y := (screenH - h - g.line) / 2

	g.xs = make([]int32, g.nx+1)
	g.ys = make([]int32, g.ny+1)

	for i := range g.xs {
		g.xs[i] = x
		x += g.size
	}
	for i := range g.ys {
		g.ys[i] = y
		y += g.size
	}
}

func (g *Grid) ScreenToGrid(pos rl.Vector2) (x, y int, in bool) {
	posX, posY := int32(pos.X), int32(pos.Y)

	quoX := (posX - g.xs[0]) / g.size
	remX := (posX - g.xs[0]) % g.size
	quoY := (posY - g.ys[0]) / g.size
	remY := (posY - g.ys[0]) % g.size

	in = quoX >= 0 && quoX < g.nx && remX >= g.line &&
		quoY >= 0 && quoY < g.ny && remY >= g.line
	if !in {
		return
	}

	x = int(quoX)
	y = int(quoY)

	return
}

func (g *Grid) DrawGrid(color color.RGBA) {
	w, h := g.WH()

	for _, posX := range g.xs {
		rl.DrawRectangle(posX, g.ys[0], g.line, h, color)
	}

	for _, posY := range g.ys {
		rl.DrawRectangle(g.xs[0], posY, w, g.line, color)
	}
}

func (g *Grid) DrawTile(x, y int, color color.RGBA) {
	posX, posY := g.xs[x]+g.line, g.ys[y]+g.line
	rectSize := g.size - g.line
	rl.DrawRectangle(posX, posY, rectSize, rectSize, color)
}
