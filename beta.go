package main

import "fmt"

func Beta() {
    fmt.Println("Beta calls Alpha")
    Alpha()
}

func Gamma() {
    fmt.Println("Gamma does not call anyone")
}