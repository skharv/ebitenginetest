package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type EntityManager struct {
}

var (
	entities []Character
)

func (e *EntityManager) Init() {
	var player Character
	player.Init(screenWidth/2, screenHeight/2, "images/babby.png")

	entities = append(entities, player)
}

func (e *EntityManager) Update() {
	for i := range entities {
		entities[i].Update()
	}
}

func (e *EntityManager) Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions) {
	for i := range entities {
		entities[i].Draw(screen, options)
	}

}

func (e *EntityManager) ReadInputs() {
	for i := range entities {
		entities[i].ReadInputs()
	}
}
