package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type EntityManager struct {
}

var (
	characters []Character
	players    []PlayerCharacter
)

func (e *EntityManager) Init() {
	var npc Character
	npc.Init(screenWidth/2, screenHeight/2, "images/face.png")

	characters = append(characters, npc)

	var player PlayerCharacter
	player.Init(screenWidth/2, screenHeight/2, "images/babby.png")

	players = append(players, player)
}

func (e *EntityManager) ReadInputs() {
	for i := range characters {
		characters[i].ReadInputs()
	}

	for i := range players {
		players[i].ReadInputs()
	}
}

func (e *EntityManager) Update() {
	for i := range characters {
		characters[i].Update()
	}

	for i := range players {
		players[i].Update()
	}
}

func (e *EntityManager) Draw(screen *ebiten.Image) {
	for i := range characters {
		characters[i].Draw(screen)
	}

	for i := range players {
		players[i].Draw(screen)
	}
}
