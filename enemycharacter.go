package main

import (
	"math/rand"

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

func (e *EnemyCharacter) Update() {

	var xRand, yRand int

	xRand = rand.Intn(5) - 2
	yRand = rand.Intn(5) - 2

	e.character.posX += float64(xRand)
	e.character.posY += float64(yRand)

	e.character.Update()
}

func (e *EnemyCharacter) Draw(screen *ebiten.Image) {
	e.character.Draw(screen)
}
