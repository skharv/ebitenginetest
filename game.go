package main

import (
	_ "image/png"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

var (
	newTime, oldTime int64
)

type Game struct {
	sceneManager *SceneManager
}

func (g *Game) Update() error {
	oldTime = newTime
	newTime = time.Now().UnixNano()
	deltaTime := float64((newTime-oldTime)/1000000) * 0.001

	if g.sceneManager == nil {
		g.sceneManager = &SceneManager{}
		g.sceneManager.GoTo(&TitleScene{}, 20)
	}

	g.sceneManager.ReadInput()
	if err := g.sceneManager.Update(deltaTime); err != nil {
		return err
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Init() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Spaceshit Game")
}
