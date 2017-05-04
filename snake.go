package main

import (
    tl "github.com/JoelOtter/termloop"
)

type Direction uint8

const (
    DirectionUp Direction = iota
    DirectionRight
    DirectionDown
    DirectionLeft
)

type Snake struct {
    *tl.Entity
    frequency float64
    direction Direction
    lastMoved float64
    alive bool
}

func NewSnake(x, y int, color tl.Attr) *Snake {
    s := &Snake { tl.NewEntity(x, y, 1, 1), 0.1, DirectionRight, 0.0, false }
    s.SetCell(0, 0, &tl.Cell{Bg: color, Ch: 's'})
    return s
}

func (s *Snake) MoveManual(event tl.Event) {
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

func (s *Snake) Move() {
    if !s.alive {
        return
    }
    offsets := map[Direction]([2]int) {
        DirectionUp:    { 0, 1 },
        DirectionDown:  { 0, -1 },
        DirectionLeft:  { -1, 0 },
        DirectionRight: { 1, 0 },
    }
    x, y := s.Position()
    ofs := offsets[s.direction]
    s.SetPosition(x + ofs[0], y + ofs[1])
}

func (s *Snake) Tick(event tl.Event) {
    directionChars := map[rune]Direction{
        'l': DirectionRight,
        'h': DirectionLeft,
        'k': DirectionDown,
        'j': DirectionUp,
    }
    directionKeys := map[tl.Key]Direction{
        tl.KeyArrowRight: DirectionRight,
        tl.KeyArrowLeft:  DirectionLeft,
        tl.KeyArrowUp:    DirectionUp,
        tl.KeyArrowDown:  DirectionDown,
    }
    speedChars := map[rune]float64 {
        '+':  0.01,
        '-': -0.01,
    }
    if event.Type == tl.EventKey {
        dir, ok := directionKeys[event.Key]
        if ok {
            s.direction = dir
        }
        dir, ok = directionChars[event.Ch]
        if ok {
            s.direction = dir
        }
        accel, ok := speedChars[event.Ch]
        if ok {
            s.frequency += accel
        }
        if event.Ch == 's' {
            s.alive = true
        }
    }
}

func (s *Snake) Draw(screen *tl.Screen) {
    s.lastMoved += screen.TimeDelta()
    if s.lastMoved > s.frequency {
        s.Move()
        s.lastMoved -= s.frequency
    }
    s.Entity.Draw(screen)
}

func (s *Snake) Collide(other tl.Physical) {
    ox, oy := other.Position()
    ow, oh := other.Size()
    x, y := s.Position()
    if x < ox || x > ox + ow {
        s.alive = false
    }
    if x < ox || x > ox + ow || y > oy || y < oy - oh {

    }
}
