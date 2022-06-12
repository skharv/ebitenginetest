package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Character struct {
	name     string
	posX     float64
	posY     float64
	rotation float64
	width    int
	height   int
	sprite   *ebiten.Image
}

func (c *Character) Init(Name string, X, Y float64, ImageFilepath string) {
	img, _, err := ebitenutil.NewImageFromFile(ImageFilepath)
	if err != nil {
		log.Fatal(err)
	}

	c.name = Name
	c.posX = X
	c.posY = Y
	c.rotation = 0
	c.sprite = img
	c.width, c.height = c.sprite.Size()
}

func (c *Character) ReadInputs() {
}

func (c *Character) Update(deltaTime float64) {

}

func (c *Character) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	options.GeoM.Scale(1, 1)
	options.GeoM.Translate(-float64(c.width)/2, -float64(c.height)/2)
	options.GeoM.Rotate(c.rotation + (90 * 3.14159 / 180))
	options.GeoM.Translate(c.posX, c.posY)

	screen.DrawImage(c.sprite, options)
}
