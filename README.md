# struct2cmd
transform the struct in golang into a cmdline

use the reflect of golang to transform your struct into a cmdline program.


## Quick Start

```go
package main

import (
  "github.com/WinChua/struct2cmd"
  "fmt"
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
  for i:=0; i< a.comm * int(a.Time); i++ {
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
```

```bash
$ go build
$ ./hmm -h
Usage of ./hmm:
  -method string
    	method should in [Hello]
  -name string
    	Name
```

## The flag and the method

By default, the field of your struct will be the flag of the cmdline.

Additionally, an option called "method" will be added which could be used to specify the method name to be call.


## Stargazers over time

[![Stargazers over time](https://starchart.cc/WinChua/struct2cmd.svg)](https://starchart.cc/WinChua/struct2cmd)

