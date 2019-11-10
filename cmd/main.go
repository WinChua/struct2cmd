package main

import (
	"fmt"
	"github.com/WinChua/struct2cmd"
)

type A struct {
	comm int
	Name string
	Time int64
	msg  string
}

func (a *A) setUp() {
	a.comm = 4
}

func (a *A) Hello() {
	for i := 0; i < a.comm*int(a.Time); i++ {
		fmt.Println(a.Name)
	}
	a.msg = "Success"
}

func (a *A) showResult() {
	fmt.Println(a.msg)
}

func main() {
	a := &A{}
	a.setUp()
	struct2cmd.Run(a)
	a.showResult()
}
