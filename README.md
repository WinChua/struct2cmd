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
  Name string
}

func (a A) Hello() {
  fmt.Println(a.Name)
}

func main() {
  a := &A{}
  struct2cmd.Run(a)
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
