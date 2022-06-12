package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	Init(Name string, X, Y float64, ImageFilepath string)
	ReadInput()
	Update(deltaTime float64)
	Draw(screen *ebiten.Image)
}

type EntityManager struct {
	entities []Entity
}

func (e *EntityManager) Init() {

	for i := 0; i < 9; i++ {
		npc := &EnemyCharacter{}
		npc.Init("Baddie "+fmt.Sprint(i), float64(50*i), float64(50*i), "images/enemy.png")

		e.entities = append(e.entities, npc)
	}

	player := &PlayerCharacter{}
	player.Init("You", ScreenWidth/2, ScreenHeight/2, "images/player.png")

	e.entities = append(e.entities, player)
}

func (e *EntityManager) ReadInput() {
	for i := range e.entities {
		e.entities[i].ReadInput()
	}
}

func (e *EntityManager) Update(deltaTime float64) {
	for i := range e.entities {
		e.entities[i].Update(deltaTime)
	}
}

func (e *EntityManager) Draw(screen *ebiten.Image) {
	for i := range e.entities {
		e.entities[i].Draw(screen)
	}
}
