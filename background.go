package main

import (
    tl "github.com/JoelOtter/termloop"
)

type Background struct {
    *tl.Rectangle
}

func (b *Background) Draw(screen *tl.Screen) {
    b.SetSize(screen.Size())
    b.Rectangle.Draw(screen)
}

func NewBackground(color tl.Attr) *Background {
    return &Background { tl.NewRectangle(0, 0, 1, 1, color) }
}
