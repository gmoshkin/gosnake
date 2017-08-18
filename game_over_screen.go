package main

import (
    tl "github.com/gmoshkin/termloop"
)

const (
    GameOverTitleFgColor tl.Attr = tl.ColorRed
    GameOverTitleBgColor tl.Attr = tl.ColorDefault
    GameOverButtonBgColor tl.Attr = tl.ColorBlue
    GameOverButtonFgColor tl.Attr = tl.ColorWhite
    GameOverButtonTextXMargin int = 2
    GameOverButtonTextYMargin int = 1
)

//////////////////////////////// GameOverButton ////////////////////////////////

type GameOverButton struct {
    *Button
}

func NewGameOverButton() *GameOverButton {
    return &GameOverButton {
        NewButton(
            func() { g.SetScreen(NewGameScreen()) },
            "press enter to start over",
            GameOverButtonFgColor,
            GameOverButtonBgColor,
            GameOverButtonTextXMargin,
            GameOverButtonTextYMargin,
        ),
    }
}

func (b *GameOverButton) Draw(screen *tl.Screen) {
    scrnW, scrnH := screen.Size()
    btnW, btnH := b.GetSize()
    b.SetPosition((scrnW - btnW) / 2, (scrnH - btnH) / 2 + 2)
    b.Button.Draw(screen)
}

//////////////////////////////// GameOverTitle /////////////////////////////////

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

//////////////////////////////// GameOverScreen ////////////////////////////////

func NewGameOverScreen() *tl.Screen {
    s := tl.NewScreen()
    s.SetFps(30)
    s.AddEntity(NewGameOverTitle())
    s.AddEntity(NewGameOverButton())
    return s
}
