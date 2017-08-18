package main

import (
    tl "github.com/gmoshkin/termloop"
)

type Callback func ()

type Button struct {
    *tl.Rectangle
    callback Callback
    text *tl.Text
    textOfsX int
    textOfsY int
    wasPressed bool
}

func NewButton(callback Callback, text string, fgColor, bgColor tl.Attr,
               xMargin, yMargin int) *Button {
    return &Button {
        tl.NewRectangle(1, 1, len(text) + 2 * xMargin, 1 + 2 * yMargin, bgColor),
        callback,
        tl.NewText(1, 1, text, fgColor, bgColor),
        xMargin,
        yMargin,
        false,
    }
}

func (b *Button) GetSize() (x, y int) {
    return b.Rectangle.Size()
}

func (b *Button) SetPosition(x, y int) {
    b.Rectangle.SetPosition(x, y)
    b.text.SetPosition(x + b.textOfsX, y + b.textOfsY)
}

func (b *Button) Tick(event tl.Event) {
    switch event.Type {
    case tl.EventKey:
        if event.Key == tl.KeyEnter {
            b.wasPressed = true
        }
    }
}

func (b *Button) Draw(screen *tl.Screen) {
    b.Rectangle.Draw(screen)
    b.text.Draw(screen)
    if b.wasPressed {
        b.callback()
    }
}
