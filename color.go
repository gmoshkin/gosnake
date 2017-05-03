package main

import (
    "fmt"
    "encoding/hex"
    "github.com/nsf/termbox-go"
)

type Color interface {
    GetTermColor() int
}

type TermColor int

func (clr TermColor) GetTermColor() int {
    return int(clr)
}

func GetTermFg(clr int) string {
    if clr > 0 {
        return fmt.Sprintf("38;5;%d", clr)
    } else {
        return "90"
    }
}

func GetTermBg(clr int) string {
    if clr > 0 {
        return fmt.Sprintf("48;5;%d", clr)
    } else {
        return "49"
    }
}

func ToTerm(clr uint8) int {
    ratio := 6 / 256.0
    return int(float64(clr) * ratio)
}

type RGBColor struct {
    red uint8
    green uint8
    blue uint8
    alpha uint8
}

func (clr RGBColor) GetTermColor() int {
    if clr.alpha < 255 / 2 {
        return -1
    } else {
        r := ToTerm(clr.red)
        g := ToTerm(clr.green)
        b := ToTerm(clr.blue)
        return int(16 + 36 * r + 6 * g + b)
    }
}

func Hex2RGB(hex_color string) RGBColor {
    start := 0
    if hex_color[0] == '#' {
        start = 1
    }
    vals, _ := hex.DecodeString(hex_color[start:])
    a := uint8(255)
    if len(vals) > 3 {
        a = vals[3]
    }
    rgb := RGBColor{ vals[0], vals[1], vals[2], a }
    return rgb
}

func GetTermboxAttributes(top *Color, bottom *Color) (termbox.Attribute, termbox.Attribute) {
    fg, bg := 0, 0
    if top == nil {
        bg = -1
    } else {
        bg = (*top).GetTermColor()
    }
    if bottom == nil {
        fg = -1
    } else {
        fg = (*bottom).GetTermColor()
    }
    return termbox.Attribute(fg), termbox.Attribute(bg)
}
