package main

import (
    tl "github.com/gmoshkin/termloop"
)

const (
    StartButtonBgColor tl.Attr = tl.ColorGreen
    StartButtonFgColor tl.Attr = tl.ColorWhite
    StartButtonTextXMargin int = 2
    StartButtonTextYMargin int = 1
)

///////////////////////////////// StartButton //////////////////////////////////

type StartButton struct {
    *Button
}

func NewStartButton() *StartButton {
    return &StartButton {
        NewButton(
            func() { g.SetScreen(NewGameScreen()) },
            "PRESS ENTER",
            StartButtonFgColor,
            StartButtonBgColor,
            StartButtonTextXMargin,
            StartButtonTextYMargin,
        ),
    }
}

func (b *StartButton) Draw(screen *tl.Screen) {
    scrnW, scrnH := screen.Size()
    btnW, btnH := b.GetSize()
    b.SetPosition((scrnW - btnW) / 2, (scrnH - btnH) / 2)
    b.Button.Draw(screen)
}

///////////////////////////////// StartScreen //////////////////////////////////

func NewStartScreen() *tl.Screen {
    s := tl.NewScreen()
    s.SetFps(30)
    button := NewStartButton()
    s.AddEntity(button)
    return s
}
