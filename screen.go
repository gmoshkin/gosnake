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
    pixels map[Position]int
    width int
    height int
}

func (screen Screen) GetWidth() int {
    return screen.width
}

func (screen Screen) GetHeight() int {
    return screen.height
}

func (screen Screen) GetPixel(row int, col int) int {
    return screen.pixels[Position{row, col}]
}

func (screen Screen) Display() {
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
    for i := 0; i < screen.GetHeight() / 2; i++ {
        row1, row2 := 2 * i, 2 * i + 1
        for col := 0; col < screen.GetWidth(); col++ {
            bg := screen.GetPixel(row1, col)
            fg := screen.GetPixel(row2, col)
            DisplayDixel(i, col, fg, bg)
        }
    }
    termbox.Flush()
}

func MakeScreen(width int, height int) *Screen {
    screen := Screen{}
    screen.width = width
    screen.height = height
    screen.pixels = make(map[Position]int)
    return &screen
}

func (screen Screen) SetCell(row int, col int, clr int) {
    screen.pixels[Position{row, col}] = clr
}

func DisplayDixel(row int, col int, fg int, bg int) {
    fgA, bgA := termbox.ColorDefault, termbox.ColorDefault
    if fg == 0 {
        fgA = termbox.Attribute(9)
    } else {
        fgA = termbox.Attribute(fg)
    }
    if bg != 0 {
        bgA = termbox.Attribute(bg)
    }
    termbox.SetCell(col, row, littleSquare, fgA, bgA)
}
