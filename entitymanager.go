package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type EntityManager struct {
}

var (
	characters []EnemyCharacter
	players    []PlayerCharacter
)

func (e *EntityManager) Init() {

	for i := 0; i < 9; i++ {
		var npc EnemyCharacter
		npc.Init("Face "+fmt.Sprint(i), float64(50*i), float64(50*i), "images/enemy.png")

		characters = append(characters, npc)
	}

	var player PlayerCharacter
	player.Init("Cute Baby", screenWidth/2, screenHeight/2, "images/player.png")

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

func (e *EntityManager) Update(deltaTime float64) {
	for i := range characters {
		characters[i].Update(deltaTime)
	}

	for i := range players {
		players[i].Update(deltaTime)
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
