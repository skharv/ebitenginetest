package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PlayerCharacter struct {
	obj                   GameObject
	up, down, left, right bool
}

func (p *PlayerCharacter) Init(Name string, X, Y float64, ImageFilepath string) {
	p.obj.Init(Name, X, Y, ImageFilepath)
	p.obj.rotSpeed = float64(5)
	p.obj.speed = float64(50)
}

func (p *PlayerCharacter) ReadInput() {
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

	p.obj.ReadInput()
}

func (p *PlayerCharacter) Update(deltaTime float64) {

	if p.left {
		p.obj.rotation = p.obj.rotation - p.obj.rotSpeed*deltaTime
	}
	if p.right {
		p.obj.rotation = p.obj.rotation + p.obj.rotSpeed*deltaTime
	}
	if p.up {
		p.obj.velocityX = p.obj.speed * math.Cos(p.obj.rotation) * deltaTime
		p.obj.velocityY = p.obj.speed * math.Sin(p.obj.rotation) * deltaTime
	}
	if p.down {
		p.obj.velocityX = -p.obj.speed * math.Cos(p.obj.rotation) * deltaTime
		p.obj.velocityY = -p.obj.speed * math.Sin(p.obj.rotation) * deltaTime
	}

	p.obj.posX += p.obj.velocityX
	p.obj.posY += p.obj.velocityY

	p.obj.Update(deltaTime)
}

func (p *PlayerCharacter) Draw(screen *ebiten.Image) {
	p.obj.Draw(screen)
}
