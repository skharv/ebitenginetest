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
	player.Init("Cute Baby", ScreenWidth/2, ScreenHeight/2, "images/player.png")

	players = append(players, player)
}

func (e *EntityManager) ReadInput() {
	for i := range characters {
		characters[i].ReadInput()
	}

	for i := range players {
		players[i].ReadInput()
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
