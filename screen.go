package main

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
    // return Hex2RGB("#ffffff")
    return screen.pixels[Position{row, col}]
}

func (screen Screen) Display() {
    for i := uint(0); i < screen.GetHeight() / 2; i++ {
        row1, row2 := 2 * i, 2 * i + 1
        for col := uint(0); col < screen.GetWidth(); col++ {
            topPtr := screen.GetPixel(row1, col)
            bottomPtr := screen.GetPixel(row2, col)
            DisplayDixel(topPtr, bottomPtr)
        }
        fmt.Println()
    }
}

func MakeScreen(width uint, height uint) *Screen {
    screen := Screen{}
    screen.width = width
    screen.height = height
    screen.pixels = make(map[Position]*Color)
    return &screen
}

func DisplayDixel(top *Color, bottom *Color) {
    var fg, bg string
    if top == nil {
        bg = GetTermBg(-1)
    } else {
        bg = GetTermBg((*top).GetTermColor())
    }
    if bottom == nil {
        fg = GetTermFg(-1)
    } else {
        fg = GetTermFg((*bottom).GetTermColor())
    }
    fmt.Printf("\033[%sm\033[%sm▄\033[0m", bg, fg)
}

