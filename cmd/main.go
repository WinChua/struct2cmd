package main

import (
	"fmt"
	"github.com/WinChua/struct2cmd"
)

type A struct {
	Name string `default:"a"`
}

func (a A) Hello() {
	fmt.Println(a.Name)
}

func (a *A) H() {
	fmt.Println(a.Name)
}
func (a A) World(answer int) {
	fmt.Println(answer)
}

func main() {
	a := A{}
	struct2cmd.Run(&a)
}
