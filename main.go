package main

import "github.com/hajimehoshi/ebiten/v2"

func main() {
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
