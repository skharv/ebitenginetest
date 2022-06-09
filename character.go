package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Character struct {
	posX                  float64
	posY                  float64
	width                 int
	height                int
	sprite                *ebiten.Image
	up, down, left, right bool
}

func (c *Character) Init(x, y float64, filepath string) {

	img, _, err := ebitenutil.NewImageFromFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	c.posX = x
	c.posY = y
	c.sprite = img
	c.width, c.height = c.sprite.Size()
}

func (c *Character) Update() {

	if c.up {
		c.posY--
	}
	if c.down {
		c.posY++
	}
	if c.left {
		c.posX--
	}
	if c.right {
		c.posX++
	}
}

func (c *Character) ReadInputs() {
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		c.up = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyW) {
		c.up = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		c.left = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyA) {
		c.left = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		c.down = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyS) {
		c.down = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		c.right = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyD) {
		c.right = false
	}
}

func (c *Character) Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions) {
	options.GeoM.Scale(1, 1)
	options.GeoM.Translate(c.posX-float64(c.width/2), c.posY-float64(c.height/2))

	screen.DrawImage(c.sprite, options)
}
