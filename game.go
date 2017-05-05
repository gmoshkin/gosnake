package main

import (
    tl "github.com/JoelOtter/termloop"
)

func NewLevel(bg tl.Attr) *tl.BaseLevel {
    return tl.NewBaseLevel(tl.Cell {
        Bg: bg,
    })
}

func main() {
    g := tl.NewGame()
    s := g.Screen()
    s.SetFps(60)
    g.SetDebugOn(true)
    s.EnablePixelMode()
    l := NewLevel(tl.ColorBlue)
    l.AddEntity(NewField(3, tl.ColorBlack))
    l.AddEntity(NewSnake(10, 10, tl.ColorWhite))
    s.SetLevel(l)
    g.Start()
}
