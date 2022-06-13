package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tinne26/etxt"
)

type OptionsScene struct {
	esc         bool
	txtRenderer *etxt.Renderer
}

func (o *OptionsScene) Init() {
	o.esc = false

	fontLib := etxt.NewFontLibrary()

	_, _, err := fontLib.ParseDirFonts("fonts")
	if err != nil {
		log.Fatal(err)
	}

	if !fontLib.HasFont("Blaka Regular") {
		log.Fatal("missing font Blaka-Regular.ttf")
	}

	o.txtRenderer = etxt.NewStdRenderer()
	glyphsCache := etxt.NewDefaultCache(10 * 1024 * 1024) // 10MB
	o.txtRenderer.SetCacheHandler(glyphsCache.NewHandler())
	o.txtRenderer.SetFont(fontLib.GetFont("Blaka Regular"))
	o.txtRenderer.SetAlign(etxt.YCenter, etxt.XCenter)
	o.txtRenderer.SetSizePx(24)
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

	o.txtRenderer.SetTarget(screen)
	o.txtRenderer.SetColor(color.RGBA{255, 255, 255, 255})
	o.txtRenderer.Draw("Here is where the options would go\n(If there was any)\nESC to return to main menu", ScreenWidth/2, ScreenHeight/2)
}
