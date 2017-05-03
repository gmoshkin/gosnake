package main

import (
    "github.com/nsf/termbox-go"
)

const littleSquare rune = '▄'

type Position struct {
    row uint
    col uint
}

type Screen struct {
    pixels map[Position]*Color
    width uint
    height uint
}

func (screen Screen) GetWidth() uint {
    return screen.width
}

func (screen Screen) GetHeight() uint {
    return screen.height
}

func (screen Screen) GetPixel(row uint, col uint) *Color {
    return screen.pixels[Position{row, col}]
}

func (screen Screen) Display() {
    for i := uint(0); i < screen.GetHeight() / 2; i++ {
        row1, row2 := 2 * i, 2 * i + 1
        for col := uint(0); col < screen.GetWidth(); col++ {
            topPtr := screen.GetPixel(row1, col)
            bottomPtr := screen.GetPixel(row2, col)
            fg, bg := GetTermboxAttributes(topPtr, bottomPtr)
            termbox.SetCell(int(i), int(col), littleSquare, fg, bg)
        }
    }
}

func MakeScreen(width uint, height uint) *Screen {
    screen := Screen{}
    screen.width = width
    screen.height = height
    screen.pixels = make(map[Position]*Color)
    return &screen
}
