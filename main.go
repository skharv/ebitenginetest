package main

import (
	_ "image/png"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
	sampleRate   = 32000
	refreshRate  = 60
)

var (
	eManager         EntityManager
	frameCount       int
	newTime, oldTime int64
)

type Game struct {
}

func (g *Game) Update() error {
	frameCount++
	newTime = time.Now().UnixNano()
	deltaTime := float64((newTime-oldTime)/1000000) * 0.001

	//currentTime := time.Duration(frameCount) * time.Second / refreshRate

	eManager.ReadInputs()
	eManager.Update(deltaTime)

	oldTime = newTime
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	eManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	frameCount = 0
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Isn't he cute?")

	eManager.Init()

	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
