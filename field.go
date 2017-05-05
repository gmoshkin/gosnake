package main

import (
    tl "github.com/JoelOtter/termloop"
)

type Field struct {
    *tl.Rectangle
    borderWidth int
}

func (f *Field) Draw(screen *tl.Screen) {
    w, h := screen.Size()
    f.SetSize(w - 2 * f.borderWidth, h - 2 * f.borderWidth)
    f.Rectangle.Draw(screen)
}

func NewField(bw int, color tl.Attr) *Field {
    return &Field { tl.NewRectangle(bw, bw, 1, 1, color), bw }
}
