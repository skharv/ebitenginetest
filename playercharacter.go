package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PlayerCharacter struct {
	character             Character
	up, down, left, right bool
}

func (p *PlayerCharacter) Init(x, y float64, filepath string) {
	p.character.Init(x, y, filepath)
}

func (p *PlayerCharacter) ReadInputs() {
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		p.up = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyW) {
		p.up = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		p.left = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyA) {
		p.left = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		p.down = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyS) {
		p.down = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		p.right = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyD) {
		p.right = false
	}

	p.character.ReadInputs()
}

func (p *PlayerCharacter) Update() {
	if p.up {
		p.character.posY--
	}
	if p.down {
		p.character.posY++
	}
	if p.left {
		p.character.posX--
	}
	if p.right {
		p.character.posX++
	}

	p.character.Update()
}

func (p *PlayerCharacter) Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions) {
	p.character.Draw(screen, options)
}
