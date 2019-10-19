package main

import (
    "fmt"
    "github.com/WinChua/struct2cmd"
)


type A struct {
    Name string
}

func (a A) Hello() {
    fmt.Println(a.Name)
}

func main() {
    a := A{}
    struct2cmd.Run(&a)
}
