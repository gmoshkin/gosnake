package main

import (
    tl "github.com/JoelOtter/termloop"
)

var g *tl.Game

func main() {
    g = tl.NewGame()
    g.SetDebugOn(true)
    s := g.Screen()
    s.SetFps(60)
    s.EnablePixelMode()
    l := NewSnakeLevel()
    s.SetLevel(l)
    g.Start()
}
