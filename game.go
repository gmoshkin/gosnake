package main

import (
    tl "github.com/gmoshkin/termloop"
)

var g *tl.Game

func NewGameScreen() *tl.Screen {
    s := tl.NewScreen()
    s.SetFps(60)
    s.EnablePixelMode()
    l := NewSnakeLevel()
    s.SetLevel(l)
    return s
}

func NewGame() *tl.Game {
    g = tl.NewGame()
    g.SetScreen(NewStartScreen())
    return g
}
