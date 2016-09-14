package main

import (
  "flag"
  "fmt"
)

func main() {
  var cmd string

  flag.StringVar(&cmd, "cmd", cmd, `cmd can be either "yo" or "cya"`)
  flag.Parse()

  switch cmd {
  case "yo":
    fmt.Println("What's up?!  :D ")
  case "cya":
    fmt.Println(" :( ")
  default:
    flag.Usage()
  }
}
