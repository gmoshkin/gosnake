package main

import (
    tl "github.com/JoelOtter/termloop"
    "container/list"
)

type Direction uint8

const (
    DirectionUp Direction = iota
    DirectionRight
    DirectionDown
    DirectionLeft
)

const (
    snakeAcceleration float64 = 0.04
    defaultFrequency float64 = 0.5
    defaultDirection Direction = DirectionRight
    defaultLastMoved float64 = 0.0
    defaultAlive bool = true
    defaultTailLength int = 3
    defaultTailColor tl.Attr = tl.ColorYellow
)

type Node struct {
    x, y int
}

type Tail struct {
    cells *list.List
    cell tl.Cell
}

func NewTail(x, y int, count int, color tl.Attr) *Tail {
    t := &Tail { list.New(), tl.Cell { Bg: color } }
    for i := 0; i < count; i++ {
        t.cells.PushFront(Node { x + i, y })
    }
    return t
}

func (t *Tail) Move(x, y int) {
    t.cells.Remove(t.cells.Back())
    t.cells.PushFront(Node { x, y })
}

func (t *Tail) Draw(s *tl.Screen) {
    for e := t.cells.Front(); e != nil; e = e.Next() {
        p, ok := e.Value.(Node)
        if ok {
            s.RenderCell(p.x, p.y, &t.cell)
        }
    }
}

type Snake struct {
    *tl.Entity
    tail *Tail
    frequency float64
    direction Direction
    lastMoved float64
    alive bool
}

func NewSnake(x, y int, color tl.Attr) *Snake {
    s := &Snake {
        Entity:     tl.NewEntity(x + defaultTailLength, y, 1, 1),
        tail:       NewTail(x, y, defaultTailLength, defaultTailColor),
        frequency:  defaultFrequency,
        direction:  defaultDirection,
        lastMoved:  defaultLastMoved,
        alive:      defaultAlive,
    }
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
    s.tail.Move(s.Position())
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
    s.tail.Draw(screen)
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
