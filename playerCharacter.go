package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PlayerCharacter struct {
	obj                                 GameObject
	up, down, left, right, moving, flip bool
	count                               int
}

const (
	scarfOX     = 0
	scarfOY     = 32
	scarfFrames = 5

	manOX     = 0
	manOY     = 0
	manOFlipY = 16
	manFrames = 4
)

var (
	animSpeed = 10
)

func (p *PlayerCharacter) Init(Name string, X, Y float64, FrameWidth, FrameHeight int, ImageFilepath string) {
	p.obj.Init(Name, X, Y, FrameWidth, FrameHeight, ImageFilepath)
	p.obj.speed = 50
}

func (p *PlayerCharacter) ReadInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		p.up = true
		p.moving = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyW) {
		p.up = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		p.left = true
		p.moving = true
		p.flip = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyA) {
		p.left = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		p.down = true
		p.moving = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyS) {
		p.down = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		p.right = true
		p.moving = true
		p.flip = false
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyD) {
		p.right = false
	}

	if !p.up && !p.down && !p.left && !p.right {
		p.moving = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		p.flip = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		p.flip = false
	}

	p.obj.ReadInput()
}

func (p *PlayerCharacter) Update(deltaTime float64) {
	p.count++

	if p.left {
		p.obj.velocityX -= p.obj.speed * deltaTime
	}
	if p.right {
		p.obj.velocityX += p.obj.speed * deltaTime
	}
	if p.up {
		p.obj.velocityY -= p.obj.speed * deltaTime
	}
	if p.down {
		p.obj.velocityY += p.obj.speed * deltaTime
	}

	p.obj.posX += p.obj.velocityX
	p.obj.posY += p.obj.velocityY

	p.obj.velocityX = 0
	p.obj.velocityY = 0

	p.obj.Update(deltaTime)
}

func (p *PlayerCharacter) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	options.GeoM.Scale(p.obj.scaleX, p.obj.scaleY)
	options.GeoM.Translate((-float64(p.obj.width)*p.obj.scaleX)/2, (-float64(p.obj.height)*p.obj.scaleY)/2)
	options.GeoM.Translate(p.obj.posX, p.obj.posY)

	if p.flip {
		if p.moving {
			i := (p.count / animSpeed) % manFrames
			mx, my := manOX+i*p.obj.width, manOFlipY
			screen.DrawImage(p.obj.sprite.SubImage(image.Rect(mx, my, mx+p.obj.width, my+p.obj.height)).(*ebiten.Image), options)
		} else {
			screen.DrawImage(p.obj.sprite.SubImage(image.Rect(manOX, manOFlipY, manOX+p.obj.width, manOFlipY+p.obj.height)).(*ebiten.Image), options)
		}
	} else {
		if p.moving {
			i := (p.count / animSpeed) % manFrames
			mx, my := manOX+i*p.obj.width, manOY
			screen.DrawImage(p.obj.sprite.SubImage(image.Rect(mx, my, mx+p.obj.width, my+p.obj.height)).(*ebiten.Image), options)
		} else {
			screen.DrawImage(p.obj.sprite.SubImage(image.Rect(manOX, manOY, manOX+p.obj.width, manOY+p.obj.height)).(*ebiten.Image), options)
		}
	}

	j := (p.count / animSpeed) % scarfFrames
	sx, sy := scarfOX+j*p.obj.width, scarfOY
	screen.DrawImage(p.obj.sprite.SubImage(image.Rect(sx, sy, sx+p.obj.width, sy+p.obj.height)).(*ebiten.Image), options)

}
