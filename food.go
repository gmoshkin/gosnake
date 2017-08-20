package main

import (
    tl "github.com/gmoshkin/termloop"
    "container/list"
)

const (
    FoodColor tl.Attr = tl.ColorMagenta
    FoodFrequency float64 = 8.0
)

///////////////////////////////////// Food /////////////////////////////////////

type Food struct {
    *tl.Entity
}

func NewFood(x, y int) *Food {
    f := &Food { tl.NewEntity(x, y, 1, 1) }
    f.SetCell(0, 0, &tl.Cell { Bg: FoodColor, Ch: 'f' })
    return f
}

///////////////////////////////// FoodManager //////////////////////////////////

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
