package main

import (
    tl "github.com/gmoshkin/termloop"
)

const (
    GameOverTitleFgColor tl.Attr = tl.ColorRed
    GameOverTitleBgColor tl.Attr = tl.ColorDefault
    GameOverButtonBgColor tl.Attr = tl.ColorGreen
    GameOverButtonFgColor tl.Attr = tl.ColorWhite
    GameOverButtonTextXMargin int = 2
    GameOverButtonTextYMargin int = 1
)

type GameOverButton struct {
    *tl.Rectangle
    callback Callback
    text *tl.Text
    wasPressed bool
}

func NewGameOverButton(callback Callback, text string) *GameOverButton {
    return &GameOverButton {
        tl.NewRectangle(1, 1, 1, 1, GameOverButtonBgColor),
        callback,
        tl.NewText(1, 1, text, GameOverButtonFgColor, GameOverButtonBgColor),
        false,
    }
}

func (b *GameOverButton) Tick(event tl.Event) {
    switch event.Type {
    case tl.EventKey:
        if event.Key == tl.KeyEnter {
            b.wasPressed = true
        }
    }
}

func (b *GameOverButton) Draw(screen *tl.Screen) {
    scrnW, scrnH := screen.Size()
    txtW, txtH := b.text.Size()
    btnW := txtW + GameOverButtonTextXMargin * 2
    btnH := txtH + GameOverButtonTextYMargin * 2
    btnX := (scrnW - btnW) / 2
    btnY := (scrnH - btnH) / 2 + 2
    b.Rectangle.SetSize(btnW, btnH)
    b.Rectangle.SetPosition(btnX, btnY)
    b.Rectangle.Draw(screen)
    b.text.SetPosition(btnX + GameOverButtonTextXMargin, btnY + GameOverButtonTextYMargin)
    b.text.Draw(screen)
    if b.wasPressed {
        b.callback()
    }
}

type GameOverTitle struct {
    *tl.Text
}

func (t *GameOverTitle) Draw(screen *tl.Screen) {
    scrnW, scrnH := screen.Size()
    txtW, txtH := t.Text.Size()
    t.Text.SetPosition((scrnW - txtW) / 2, (scrnH - txtH) / 2)
    t.Text.Draw(screen)
}

func NewGameOverTitle() *GameOverTitle {
    return &GameOverTitle {
        tl.NewText(1, 1, "GAME OVER", GameOverTitleFgColor, GameOverTitleBgColor),
    }
}

func NewGameOverScreen() *tl.Screen {
    s := tl.NewScreen()
    s.SetFps(30)
    s.AddEntity(NewGameOverTitle())
    s.AddEntity(NewGameOverButton(func () {
        g.SetScreen(NewGameScreen())
    },"Press Enter to start over"))
    return s
}
