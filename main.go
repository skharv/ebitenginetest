package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
	sampleRate   = 32000
)

var (
	eManager EntityManager
)

type Game struct {
}

func (g *Game) Update() error {
	eManager.ReadInputs()
	eManager.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	eManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Isn't he cute?")

	eManager.Init()

	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
