package main

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type TitleScene struct {
	play    bool
	options bool
	esc     bool
}

func (t *TitleScene) Init() {
	t.play = false
	t.options = false
	t.esc = false
}

func (t *TitleScene) ReadInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		t.play = true
	} else {
		t.play = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyO) {
		t.options = true
	} else {
		t.options = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		t.esc = true
	} else {
		t.esc = false
	}
}

func (t *TitleScene) Update(state *GameState, deltaTime float64) error {
	if t.play {
		g := &GameScene{}
		state.SceneManager.GoTo(g, 50)
	}
	if t.options {
		o := &OptionsScene{}
		state.SceneManager.GoTo(o, 10)
	}
	if t.esc {
		os.Exit(0)
	}
	return nil
}

func (t *TitleScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 1)

	ebitenutil.DebugPrint(screen, "Press SPACE to play\nPress O for options\nPress ESC to quit")
}
