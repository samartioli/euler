package main

import (
    "os"
    "fmt"
    "github.com/samartioli/euler/problems"
)

func main() {

    arg2 := os.Args[1]

    fmt.Println("Welcome to Sam's Project Euler Solutions in GO")
    _, ok := problems.P[arg2]

    if !ok {
        fmt.Println("Does not Exist")
    } else {
        problems.P[arg2]()
    }

}
