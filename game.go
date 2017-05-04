package main

import (
    tl "github.com/JoelOtter/termloop"
)

type Field struct {
    *tl.Rectangle
}

func NewField(x, y, w, h int, color tl.Attr) *Field {
    return &Field { tl.NewRectangle(x, y, w, h, color) }
}

var g *tl.Game

func main() {
    g = tl.NewGame()
    g.Screen().SetFps(60)
    g.SetDebugOn(true)
    g.Screen().EnablePixelMode()
    l := tl.NewBaseLevel(tl.Cell{
        Bg: tl.Attr(tl.ColorBlue),
    })
    l.AddEntity(NewField(3, 3, 70, 45, tl.ColorBlack))
    l.AddEntity(NewSnake(10, 10, tl.ColorWhite))
    g.Screen().SetLevel(l)
    // g.Screen().AddEntity(tl.NewFpsText(0, 0, tl.ColorWhite, tl.ColorDefault, 0.5))
    g.Start()
}
