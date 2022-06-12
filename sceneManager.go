package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	transitionFrom = ebiten.NewImage(ScreenWidth, ScreenHeight)
	transitionTo   = ebiten.NewImage(ScreenWidth, ScreenHeight)
)

type Scene interface {
	Init()
	ReadInput()
	Update(state *GameState, deltaTime float64) error
	Draw(screen *ebiten.Image)
}

const transitionMaxCount = 50

type SceneManager struct {
	current         Scene
	next            Scene
	transitionCount int
}

type GameState struct {
	SceneManager *SceneManager
}

func (s *SceneManager) ReadInput() {
	if s.transitionCount == 0 {
		s.current.ReadInput()
		return
	}
}

func (s *SceneManager) Update(deltaTime float64) error {
	if s.transitionCount == 0 {
		return s.current.Update(&GameState{
			SceneManager: s,
		}, deltaTime)
	}

	s.transitionCount--
	if s.transitionCount > 0 {
		return nil
	}

	s.current = s.next
	s.next = nil
	return nil
}

func (s *SceneManager) Draw(screen *ebiten.Image) {
	if s.transitionCount == 0 {
		s.current.Draw(screen)
		return
	}

	transitionFrom.Clear()
	s.current.Draw(transitionFrom)

	transitionTo.Clear()
	s.next.Draw(transitionTo)

	screen.DrawImage(transitionFrom, nil)

	alpha := 1 - float64(s.transitionCount)/float64(transitionMaxCount)
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, alpha)
	screen.DrawImage(transitionTo, op)
}

func (s *SceneManager) GoTo(scene Scene, fadeTime int) {
	scene.Init()
	if fadeTime > transitionMaxCount {
		fadeTime = transitionMaxCount
	}

	if s.current == nil {
		s.current = scene
	} else {
		s.next = scene
		s.transitionCount = fadeTime
	}
}
