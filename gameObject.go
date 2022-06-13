package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameObject struct {
	name     string
	posX     float64
	posY     float64
	rotation float64
	scaleX   float64
	scaleY   float64
	width    int
	height   int
	sprite   *ebiten.Image

	speed        float64
	rotSpeed     float64
	velocityX    float64
	velocityY    float64
	acceleration float64
}

func (c *GameObject) Init(Name string, X, Y float64, ImageFilepath string) {
	img, _, err := ebitenutil.NewImageFromFile(ImageFilepath)
	if err != nil {
		log.Fatal(err)
	}
	c.scaleX = 1
	c.scaleY = 1
	c.name = Name
	c.posX = X
	c.posY = Y
	c.rotation = 0
	c.sprite = img
	c.width, c.height = c.sprite.Size()
	c.speed = 0
	c.rotSpeed = 0
	c.velocityX = 0
	c.velocityY = 0
	c.acceleration = 0
}

func (c *GameObject) ReadInput() {
}

func (c *GameObject) Update(deltaTime float64) {

}

func (c *GameObject) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	options.GeoM.Scale(c.scaleX, c.scaleY)
	options.GeoM.Translate((-float64(c.width)*c.scaleX)/2, (-float64(c.height)*c.scaleY)/2)
	options.GeoM.Rotate(c.rotation + (90 * 3.14159 / 180))
	options.GeoM.Translate(c.posX, c.posY)

	screen.DrawImage(c.sprite, options)
}
