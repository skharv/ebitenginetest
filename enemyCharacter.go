package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type EnemyCharacter struct {
	obj GameObject
}

func (e *EnemyCharacter) Init(Name string, X, Y float64, FrameWidth, FrameHeight int, ImageFilepath string) {
	e.obj.Init(Name, X, Y, FrameWidth, FrameHeight, ImageFilepath)
}

func (e *EnemyCharacter) ReadInput() {
	e.obj.ReadInput()
}

func (e *EnemyCharacter) Update(deltaTime float64) {
	e.obj.Update(deltaTime)
}

func (e *EnemyCharacter) Draw(screen *ebiten.Image) {
	e.obj.Draw(screen)
}
