package main

import (
    tl "github.com/JoelOtter/termloop"
)

func NewLevel(bg tl.Attr) *tl.BaseLevel {
    return tl.NewBaseLevel(tl.Cell { Bg: bg })
}

func main() {
    g := tl.NewGame()
    g.SetDebugOn(true)
    s := g.Screen()
    s.SetFps(30)
    s.EnablePixelMode()
    l := NewLevel(tl.ColorBlue)
    l.AddEntity(NewBackground(tl.ColorGreen))
    l.AddEntity(NewField(3, tl.ColorBlack))
    l.AddEntity(NewSnake(10, 10, tl.ColorWhite))
    s.SetLevel(l)
    g.Start()
}
