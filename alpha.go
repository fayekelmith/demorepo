package main

import "fmt"

func Alpha() {
    fmt.Println("Alpha calls Beta and Gamma")
    Beta()
    Gamma()
}