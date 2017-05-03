package main

import (
    "github.com/nsf/termbox-go"
)

const littleSquare rune = '▄'

type Position struct {
    row int
    col int
}

type Screen struct {
    pixels map[Position]*Color
    width int
    height int
}

func (screen Screen) GetWidth() int {
    return screen.width
}

func (screen Screen) GetHeight() int {
    return screen.height
}

func (screen Screen) GetPixel(row int, col int) *Color {
    return screen.pixels[Position{row, col}]
}

func (screen Screen) Display() {
    for i := int(0); i < screen.GetHeight() / 2; i++ {
        row1, row2 := 2 * i, 2 * i + 1
        for col := int(0); col < screen.GetWidth(); col++ {
            topPtr := screen.GetPixel(row1, col)
            bottomPtr := screen.GetPixel(row2, col)
            fg, bg := GetTermboxAttributes(topPtr, bottomPtr)
            termbox.SetCell(int(i), int(col), littleSquare, fg, bg)
        }
    }
}

func MakeScreen(width int, height int) *Screen {
    screen := Screen{}
    screen.width = width
    screen.height = height
    screen.pixels = make(map[Position]*Color)
    return &screen
}
