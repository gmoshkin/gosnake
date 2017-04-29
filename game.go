package main

import (
    "fmt"
    "encoding/hex"
    "syscall"
    "unsafe"
)

type TermSize struct {
    Row    uint16
    Col    uint16
    Xpixel uint16
    Ypixel uint16
}

func GetTermSize() *TermSize {
    ws := &TermSize{}
    retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
        uintptr(syscall.Stdin),
        uintptr(syscall.TIOCGWINSZ),
        uintptr(unsafe.Pointer(ws)))

    if int(retCode) == -1 {
        panic(errno)
    }
    return ws
}

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

func DisplayDixel(top Color, bottom Color) {
    bg := GetTermBg(top.GetTermColor())
    fg := GetTermFg(bottom.GetTermColor())
    fmt.Printf("\033[%sm\033[%sm▄\033[0m", bg, fg)
}

func main() {
    DisplayDixel(TermColor(-1), Hex2RGB("#ff0088"))
    DisplayDixel(TermColor(-1), TermColor(-1))
    DisplayDixel(Hex2RGB("#008800"), TermColor(-1))
    fmt.Println(GetTermSize())
}
