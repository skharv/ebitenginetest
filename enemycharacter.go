package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type EnemyCharacter struct {
	character Character
}

func (e *EnemyCharacter) Init(Name string, X, Y float64, ImageFilepath string) {
	e.character.Init(Name, X, Y, ImageFilepath)
}

func (e *EnemyCharacter) ReadInputs() {
	e.character.ReadInputs()
}

func (e *EnemyCharacter) Update(deltaTime float64) {
	e.character.Update(deltaTime)
}

func (e *EnemyCharacter) Draw(screen *ebiten.Image) {
	e.character.Draw(screen)
}
