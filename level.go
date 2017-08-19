package main

import (
    tl "github.com/gmoshkin/termloop"
    "math/rand"
)

const (
    LevelBorderColor tl.Attr = tl.ColorGreen
    LevelBorderWidth int = 3
    LevelColor tl.Attr = tl.ColorBlack
    MaxTries int = 100
)

////////////////////////////////// Background //////////////////////////////////

type Background struct {
    *tl.Rectangle
}

func (b *Background) Draw(screen *tl.Screen) {
    b.SetSize(screen.Size())
    b.Rectangle.Draw(screen)
}

func NewBackground() *Background {
    return &Background { tl.NewRectangle(0, 0, 1, 1, LevelBorderColor) }
}

//////////////////////////////////// Field /////////////////////////////////////

type Field struct {
    *tl.Rectangle
    borderWidth int
}

func (f *Field) Draw(screen *tl.Screen) {
    w, h := screen.Size()
    f.SetSize(w - 2 * f.borderWidth, h - 2 * f.borderWidth)
    f.Rectangle.Draw(screen)
}

func NewField() *Field {
    return &Field {
        tl.NewRectangle(LevelBorderWidth, LevelBorderWidth, 1, 1, LevelColor),
        LevelBorderWidth,
    }
}

////////////////////////////////// SnakeLevel //////////////////////////////////

type SnakeLevel struct {
    *tl.BaseLevel
    background *Background
    field *Field
    snake *Snake
    foodManager *FoodManager
}

func NewSnakeLevel() *SnakeLevel {
    l := &SnakeLevel {
        tl.NewBaseLevel(tl.Cell {}),
        NewBackground(),
        NewField(),
        NewSnake(10, 10, tl.ColorWhite),
        NewFoodManager(),
    }
    l.snake.level = l
    l.AddEntity(l.background)
    l.AddEntity(l.field)
    l.AddEntity(l.snake)
    return l
}

func randomCoordinate(start, length int) int {
    return rand.Intn(length) + start
}

func  collides(x, y, pX, pY, pW, pH int) bool {
    if x < pX || y < pY {
        return false
    }
    if x >= pX + pW || y >= pY + pH {
        return false
    }
    return true
}

func (sl *SnakeLevel) isVacant(x, y int) bool {
    for _, e := range sl.Entities {
        switch e.(type) {
            case *Field:
                continue
            case *Background:
                continue
            case tl.Physical:
                p, _ := e.(tl.Physical)
                pX, pY := p.Position()
                pW, pH := p.Size()
                return !collides(x, y, pX, pY, pW, pH)
        }
    }
    return false
}

func (sl *SnakeLevel) GetVacantPoint() (x, y int) {
    startX, startY := sl.field.Position()
    width, height := sl.field.Size()
    if width < 0 || height < 0 {
        return -1, -1
    }
    leftTries := MaxTries
    for leftTries != 0 {
        x = randomCoordinate(startX, width)
        y = randomCoordinate(startY, height)
        if sl.isVacant(x, y) {
            return x, y
        }
        leftTries--
    }
    return -1, -1
}

func (sl *SnakeLevel) Draw(screen *tl.Screen) {
    if sl.foodManager.IsTime(screen.TimeDelta()) {
        x, y := sl.GetVacantPoint()
        if x > 0 && y > 0 {
            food := sl.foodManager.MakeFood(x, y)
            sl.AddEntity(food)
        }
    }
    sl.BaseLevel.Draw(screen)
}

func (sl *SnakeLevel) FoodGone(food *Food) {
    sl.RemoveEntity(food)
}

func (sl *SnakeLevel) IsBorder(x, y int) bool {
    width, height := sl.background.Size()
    return (x < LevelBorderWidth || y < LevelBorderWidth ||
            x >= width - LevelBorderWidth || y >= height - LevelBorderWidth)
}
