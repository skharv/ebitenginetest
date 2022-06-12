package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PlayerCharacter struct {
	character                               Character
	up, down, left, right, strafeL, strafeR bool
	speed                                   float64
	acc                                     float64
}

func (p *PlayerCharacter) Init(Name string, X, Y float64, ImageFilepath string) {
	p.character.Init(Name, X, Y, ImageFilepath)
	p.acc = 5
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
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		p.strafeL = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyQ) {
		p.strafeL = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		p.strafeR = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyE) {
		p.strafeR = false
	}

	p.character.ReadInputs()
}

func (p *PlayerCharacter) Update(deltaTime float64) {
	rotSpeed := float64(5)
	strafeSpeed := float64(2)

	if p.left {
		p.character.rotation = p.character.rotation - rotSpeed*deltaTime
	}
	if p.right {
		p.character.rotation = p.character.rotation + rotSpeed*deltaTime
	}
	if p.up {
		p.speed += p.acc * deltaTime
	}
	if p.down {
		p.speed -= p.acc * deltaTime
	}
	if p.strafeL {
		p.character.posX += strafeSpeed * math.Cos(p.character.rotation-(90*3.14159/180))
		p.character.posY += strafeSpeed * math.Sin(p.character.rotation-(90*3.14159/180))
	}
	if p.strafeR {
		p.character.posX += strafeSpeed * math.Cos(p.character.rotation+(90*3.14159/180))
		p.character.posY += strafeSpeed * math.Sin(p.character.rotation+(90*3.14159/180))
	}

	p.character.posX += p.speed * math.Cos(p.character.rotation)
	p.character.posY += p.speed * math.Sin(p.character.rotation)

	p.character.Update(deltaTime)
}

func (p *PlayerCharacter) Draw(screen *ebiten.Image) {
	p.character.Draw(screen)
}
