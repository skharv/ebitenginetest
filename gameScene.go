package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameScene struct {
	eManager EntityManager
	esc      bool
}

func (g *GameScene) Init() {
	g.esc = false
	g.eManager.Init()
}

func (g *GameScene) ReadInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.esc = true
	} else {
		g.esc = false
	}
	g.eManager.ReadInput()
}

func (g *GameScene) Update(state *GameState, deltaTime float64) error {
	g.eManager.Update(deltaTime)
	if g.esc {
		t := &TitleScene{}
		state.SceneManager.GoTo(t, 20)
	}
	return nil
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 64, 255})
	g.eManager.Draw(screen)
}
