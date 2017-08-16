package main

import (
    tl "github.com/gmoshkin/termloop"
)

var g *tl.Game

type Callback func ()

type StartButton struct {
    *tl.Rectangle
    callback Callback
    wasPressed bool
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
    b.Rectangle.Draw(screen)
    if b.wasPressed {
        b.callback()
    }
}

func NewStartScreen(callback Callback) *tl.Screen {
    s := tl.NewScreen()
    s.SetFps(30)
    s.AddEntity(&StartButton {
        tl.NewRectangle(10, 10, 10, 10, tl.ColorRed),
        callback,
        false,
    })
    return s
}

func NewGameScreen() *tl.Screen {
    s := tl.NewScreen()
    s.SetFps(60)
    s.EnablePixelMode()
    l := NewSnakeLevel()
    s.SetLevel(l)
    return s
}

func main() {
    g = tl.NewGame()
    g.SetDebugOn(true)
    gameScreen := NewGameScreen()
    onStart := func () {
        g.SetScreen(gameScreen)
    }
    startScreen := NewStartScreen(onStart)
    g.SetScreen(startScreen)
    g.Start()
}
