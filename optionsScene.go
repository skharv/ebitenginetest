package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type OptionsScene struct {
	esc bool
}

func (o *OptionsScene) Init() {
	o.esc = false
}

func (o *OptionsScene) ReadInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		o.esc = true
	} else {
		o.esc = false
	}
}

func (o *OptionsScene) Update(state *GameState, deltaTime float64) error {
	if o.esc {
		t := &TitleScene{}
		state.SceneManager.GoTo(t, 10)
	}
	return nil
}

func (o *OptionsScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 1)

	ebitenutil.DebugPrint(screen, "This is where options would go (if we had any)")
}
