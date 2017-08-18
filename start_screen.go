package main

import (
    tl "github.com/gmoshkin/termloop"
)

const (
    ButtonBgColor tl.Attr = tl.ColorGreen
    ButtonFgColor tl.Attr = tl.ColorWhite
    ButtonTextXMargin int = 2
    ButtonTextYMargin int = 1
)

type Callback func ()

type StartButton struct {
    *tl.Rectangle
    callback Callback
    text *tl.Text
    wasPressed bool
}

func NewStartButton(callback Callback, text string) *StartButton {
    return &StartButton {
        tl.NewRectangle(1, 1, 1, 1, ButtonBgColor),
        callback,
        tl.NewText(1, 1, text, ButtonFgColor, ButtonBgColor),
        false,
    }
}

func (b *StartButton) Tick(event tl.Event) {
    switch event.Type {
    case tl.EventKey:
        if event.Key == tl.KeyEnter {
            b.wasPressed = true
        }
    }
}

func (b *StartButton) Draw(screen *tl.Screen) {
    scrnW, scrnH := screen.Size()
    txtW, txtH := b.text.Size()
    btnW := txtW + ButtonTextXMargin * 2
    btnH := txtH + ButtonTextYMargin * 2
    btnX := (scrnW - btnW) / 2
    btnY := (scrnH - btnH) / 2
    b.Rectangle.SetSize(btnW, btnH)
    b.Rectangle.SetPosition(btnX, btnY)
    b.Rectangle.Draw(screen)
    b.text.SetPosition(btnX + ButtonTextXMargin, btnY + ButtonTextYMargin)
    b.text.Draw(screen)
    if b.wasPressed {
        b.callback()
    }
}

func NewStartScreen(callback Callback) *tl.Screen {
    s := tl.NewScreen()
    s.SetFps(30)
    s.AddEntity(NewStartButton(callback, "Press Enter to start"))
    return s
}
