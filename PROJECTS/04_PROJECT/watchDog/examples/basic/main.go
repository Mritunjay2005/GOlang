package main

import (
    "fmt"
    "ggithub.com/Mritunjay2005/watchDog"
)

func main() {
    w := watchdog.New(watchdog.WithDebounce(100 * time.Millisecond))
    w.On(watchdog.Modified, "**/*.go", func(e watchdog.Event) {
        fmt.Printf("changed: %s\n", e.Path)
    })
    w.Start(".")
}