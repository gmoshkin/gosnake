package main

import (
    tl "github.com/JoelOtter/termloop"
    "container/list"
)

const (
    LevelBorderColor tl.Attr = tl.ColorGreen
    LevelBorderWidth int = 3
    LevelColor tl.Attr = tl.ColorBlack
    FoodColor tl.Attr = tl.ColorYellow
    FoodFrequency float64 = 5.0
)

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

type Food struct {
    *tl.Entity
}

func NewFood(x, y int) *Food {
    f := &Food { tl.NewEntity(x, y, 1, 1) }
    f.SetCell(0, 0, &tl.Cell { Bg: FoodColor, Ch: 'f' })
    return f
}

type FoodManager struct {
    foods *list.List
    updatePeriod float64
    lastUpdate float64
}

func NewFoodManager() *FoodManager {
    return &FoodManager { list.New(), FoodFrequency, FoodFrequency - 0.5 }
}

func (fm *FoodManager) MakeFood(x, y int) *Food {
    return NewFood(x, y)
}

func (fm *FoodManager) IsTime(timeDelta float64) bool {
    fm.lastUpdate += timeDelta
    if fm.lastUpdate > fm.updatePeriod {
        fm.lastUpdate -= fm.updatePeriod
        return true
    }
    return false
}

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
    l.AddEntity(l.background)
    l.AddEntity(l.field)
    l.AddEntity(l.snake)
    return l
}

func (sl *SnakeLevel) GetFreePoint() (x, y int) {
    startX, startY := sl.field.Position()
    width, height := sl.field.Size()
    return (width - startX) / 2, (height - startY) / 2 // FIXME
}

func (sl *SnakeLevel) Draw(screen *tl.Screen) {
    if sl.foodManager.IsTime(screen.TimeDelta()) {
        x, y := sl.GetFreePoint()
        food := sl.foodManager.MakeFood(x, y)
        sl.AddEntity(food)
    }
    sl.BaseLevel.Draw(screen)
}
