package problems

import (
    "fmt"
)

var P = make(map[string]func())

func P0002() {
    fmt.Println("0002")
}

func init() {
    P["P0002"] = P0002
}
