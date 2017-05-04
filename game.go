package main

import (
    tl "github.com/JoelOtter/termloop"
)

type Snake struct {
    *tl.Entity
}

func NewSnake(x, y int, color tl.Attr) *Snake {
    s := &Snake{tl.NewEntity(x, y, 1, 1)}
    s.SetCell(0, 0, &tl.Cell{Fg: color, Ch: 's'})
    return s
}

func (s *Snake) Tick(event tl.Event) {
    charActions := map[rune]([2]int){
        'l': { 1, 0 }, 'h': { -1, 0 }, 'k': { 0, -1 }, 'j': { 0, 1 },
    }
    keyActions := map[tl.Key]([2]int){
        tl.KeyArrowRight: { 1, 0 },
        tl.KeyArrowLeft: { -1, 0 },
        tl.KeyArrowUp: { 0, -1 },
        tl.KeyArrowDown: { 0, 1 },
    }
    if event.Type == tl.EventKey {
        x, y := s.Position()
        ofs, ok := keyActions[event.Key]
        if ok {
            s.SetPosition(x + ofs[0], y + ofs[1])
        }
        ofs, ok = charActions[event.Ch]
        if ok {
            s.SetPosition(x + ofs[0], y + ofs[1])
        }
    }
}

var g *tl.Game

func main() {
    g = tl.NewGame()
    g.Screen().SetFps(60)
    g.SetDebugOn(true)
    l := tl.NewBaseLevel(tl.Cell{
        Bg: tl.Attr(tl.ColorBlue),
    })
    l.AddEntity(NewSnake(1, 1, tl.ColorWhite))
    g.Screen().SetLevel(l)
    g.Screen().AddEntity(tl.NewFpsText(0, 0, tl.ColorWhite, tl.ColorDefault, 0.5))
    g.Start()
}
