package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Character struct {
	posX   float64
	posY   float64
	width  int
	height int
	sprite *ebiten.Image
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

func (c *Character) ReadInputs() {
}

func (c *Character) Update() {
}

func (c *Character) Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions) {
	options.GeoM.Scale(1, 1)
	options.GeoM.Translate(c.posX-float64(c.width/2), c.posY-float64(c.height/2))

	screen.DrawImage(c.sprite, options)
}
