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

const snakeAcceleration float64 = 0.04

type Snake struct {
    *tl.Entity
    frequency float64
    direction Direction
    lastMoved float64
    alive bool
}

func NewSnake(x, y int, color tl.Attr) *Snake {
    s := &Snake { tl.NewEntity(x, y, 1, 1), 0.1, DirectionRight, 0.0, true }
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
        DirectionUp:    { 0, -1 },
        DirectionDown:  { 0, 1 },
        DirectionLeft:  { -1, 0 },
        DirectionRight: { 1, 0 },
    }
    x, y := s.Position()
    ofs := offsets[s.direction]
    s.SetPosition(x + ofs[0], y + ofs[1])
}

func (s *Snake) SetDirection(dir Direction) {
    switch dir {
    case DirectionUp, DirectionDown:
        switch s.direction {
        case DirectionLeft, DirectionRight:
            s.direction = dir
        }
    case DirectionLeft, DirectionRight:
        switch s.direction {
        case DirectionUp, DirectionDown:
            s.direction = dir
        }
    }
}

func (s *Snake) Tick(event tl.Event) {
    if event.Type == tl.EventKey {
        switch event.Key {
        case tl.KeyArrowUp:
            s.SetDirection(DirectionUp)
        case tl.KeyArrowLeft:
            s.SetDirection(DirectionLeft)
        case tl.KeyArrowDown:
            s.SetDirection(DirectionDown)
        case tl.KeyArrowRight:
            s.SetDirection(DirectionRight)
        }
        switch event.Ch {
        case 'w', 'k':
            s.SetDirection(DirectionUp)
        case 'a', 'h':
            s.SetDirection(DirectionLeft)
        case 's', 'j':
            s.SetDirection(DirectionDown)
        case 'd', 'l':
            s.SetDirection(DirectionRight)
        case '+':
            s.frequency -= snakeAcceleration
        case '-':
            s.frequency += snakeAcceleration
        }
    }
}

func (s *Snake) Draw(screen *tl.Screen) {
    if ! s.alive {
        return
    }
    s.lastMoved += screen.TimeDelta()
    if s.lastMoved > s.frequency {
        s.Move()
        s.lastMoved -= s.frequency
    }
    s.Entity.Draw(screen)
}

func (s *Snake) Collide(other tl.Physical) {
    switch other.(type) {
    case *Field:
        s.alive = true
    case *Background:
        s.alive = false
    }
}
