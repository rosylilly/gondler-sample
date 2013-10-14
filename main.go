package main

import (
  "flag"
  "github.com/eknkc/amber"
  "os"
)

type Val struct {
  Name string
}

func main() {
  val := &Val{}
  val.Name = *flag.String("name", "rosylilly", "")

  flag.Parse()

  template, err := amber.CompileFile("template.slim", amber.DefaultOptions)

  if err != nil {
    panic(err)
  }

  fd, err := os.Create("dist.html")
  if err != nil {
    panic(err)
  }
  defer fd.Close()

  template.Execute(fd, val)
}
