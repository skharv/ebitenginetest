package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type TitleScene struct {
	space bool
}

func (t *TitleScene) Init() {
	t.space = false
}

func (t *TitleScene) ReadInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		t.space = true
	} else {
		t.space = false
	}
}

func (t *TitleScene) Update(state *GameState, deltaTime float64) error {
	if t.space {
		g := &GameScene{}
		state.SceneManager.GoTo(g)
	}
	return nil
}

func (t *TitleScene) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 1)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Press SPACE to play"))
}
