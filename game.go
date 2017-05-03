package main

import (
    "fmt"
    "github.com/nsf/termbox-go"
)

func Init() {
    err := termbox.Init()
    if err != nil {
        panic(err)
    }
    termbox.SetOutputMode(termbox.Output256)
}

func Close() {
    termbox.Close()
}

func main() {
    Init()
    cols, rows := termbox.Size()
    screen := MakeScreen(cols, rows * 2)
    screen.SetCell(6, 6, 7)
    screen.Display()
    event := termbox.PollEvent()
    Close()
    fmt.Println(event)
    fmt.Println(screen.GetWidth(), screen.GetHeight())
}
