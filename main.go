package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	player Character
)

type Game struct {
}

func (g *Game) Update() error {
	player.ReadInputs()
	player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	player.Draw(screen, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Isn't he cute?")

	player.Init(screenWidth/2, screenHeight/2, "images/babby.png")

	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
