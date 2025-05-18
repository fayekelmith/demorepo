package main

import "fmt"

func Alpha() {
    fmt.Println("Alpha calls Gamma and Delta")
    Gamma()
    Delta()
}