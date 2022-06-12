package main

import "github.com/hajimehoshi/ebiten/v2"

type GameScene struct {
	eManager EntityManager
}

func (g *GameScene) Init() {
	g.eManager.Init()
}

func (g *GameScene) ReadInput() {
	g.eManager.ReadInput()
}

func (g *GameScene) Update(state *GameState, deltaTime float64) error {
	g.eManager.Update(deltaTime)
	return nil
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	g.eManager.Draw(screen)
}
