package main

import (
    tl "github.com/JoelOtter/termloop"
    "container/list"
)

type Direction uint8
type Acceleration float64

const (
    DirectionUp Direction = iota
    DirectionRight
    DirectionDown
    DirectionLeft
    AccelerationUp Acceleration = -0.2
    AccelerationDown Acceleration = -AccelerationUp
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

type Moves struct {
    queue *list.List
}

func (m *Moves) Add(move interface{}) {
    m.queue.PushFront(move)
}

func (m *Moves) Pop() interface{} {
    move := m.queue.Back()
    if move == nil {
        return move
    }
    m.queue.Remove(move)
    return move.Value
}

type Snake struct {
    *tl.Entity
    tail *Tail
    frequency float64
    direction Direction
    lastMoved float64
    alive bool
    moves *Moves
}

func NewSnake(x, y int, color tl.Attr) *Snake {
    s := &Snake {
        Entity:     tl.NewEntity(x + defaultTailLength, y, 1, 1),
        tail:       NewTail(x, y, defaultTailLength, defaultTailColor),
        frequency:  defaultFrequency,
        direction:  defaultDirection,
        lastMoved:  defaultLastMoved,
        alive:      defaultAlive,
        moves:      &Moves { list.New() },
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

func (s *Snake) DoAction() {
    action := s.moves.Pop()
    switch action := action.(type) {
    case Direction:
        s.SetDirection(action)
    case Acceleration:
        s.frequency += float64(action)
    }
}

func (s *Snake) Move() {
    if !s.alive {
        return
    }
    s.DoAction()
    x, y := s.Position()
    s.tail.Move(x, y)
    switch s.direction {
    case DirectionUp:
        y--
    case DirectionDown:
        y++
    case DirectionLeft:
        x--
    case DirectionRight:
        x++
    }
    s.SetPosition(x, y)
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
            s.moves.Add(DirectionUp)
        case tl.KeyArrowLeft:
            s.moves.Add(DirectionLeft)
        case tl.KeyArrowDown:
            s.moves.Add(DirectionDown)
        case tl.KeyArrowRight:
            s.moves.Add(DirectionRight)
        }
        switch event.Ch {
        case 'w', 'k':
            s.moves.Add(DirectionUp)
        case 'a', 'h':
            s.moves.Add(DirectionLeft)
        case 's', 'j':
            s.moves.Add(DirectionDown)
        case 'd', 'l':
            s.moves.Add(DirectionRight)
        case '+':
            s.moves.Add(AccelerationUp)
        case '-':
            s.moves.Add(AccelerationDown)
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
