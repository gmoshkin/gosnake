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
    termbox.SetOutputMode(termbox.Output216)
}

func Close() {
    termbox.Close()
}

func main() {
    Init()
    cols, rows := termbox.Size()
    screen := MakeScreen(cols, rows * 2)
    screen.Display()
    event := termbox.PollEvent()
    Close()
    fmt.Println(event)
}
