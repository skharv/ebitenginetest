package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PlayerCharacter struct {
	animator                            Animator
	obj                                 GameObject
	up, down, left, right, moving, flip bool
	count                               int
}

func (p *PlayerCharacter) Init(Name string, X, Y float64, FrameWidth, FrameHeight int, ImageFilepath string) {
	p.obj.Init(Name, X, Y, FrameWidth, FrameHeight, ImageFilepath)

	p.animator.Init(ImageFilepath, Vector2i{16, 16}, Vector2f{4, 4}, Vector2f{ScreenWidth / 2, ScreenHeight / 2}, 10)
	p.animator.AddAnimation(Animation{
		frameCount:         4,
		frameStartPosition: Vector2i{0, 0},
		loop:               true,
	}, "runRight")
	p.animator.AddAnimation(Animation{
		frameCount:         4,
		frameStartPosition: Vector2i{0, 16},
		loop:               true,
	}, "runLeft")
	p.animator.AddAnimation(Animation{
		frameCount:         1,
		frameStartPosition: Vector2i{0, 0},
		loop:               true,
	}, "standRight")
	p.animator.AddAnimation(Animation{
		frameCount:         1,
		frameStartPosition: Vector2i{0, 16},
		loop:               true,
	}, "standLeft")

	p.animator.SetAnimation("standRight", false)
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

	if p.moving {
		if p.flip {
			if !(p.animator.IsLooping() && p.animator.IsAnimation("runLeft")) {
				p.animator.SetAnimation("runLeft", false)
			}
		} else {
			if !(p.animator.IsLooping() && p.animator.IsAnimation("runRight")) {
				p.animator.SetAnimation("runRight", false)
			}
		}
	} else {
		if p.flip {
			if !(p.animator.IsLooping() && p.animator.IsAnimation("standLeft")) {
				p.animator.SetAnimation("standLeft", false)
			}
		} else {
			if !(p.animator.IsLooping() && p.animator.IsAnimation("standRight")) {
				p.animator.SetAnimation("standRight", false)
			}
		}
	}

	p.obj.Update(deltaTime)
	p.animator.Update(deltaTime)
}

func (p *PlayerCharacter) Draw(screen *ebiten.Image) {
	p.animator.Draw(screen)
}
