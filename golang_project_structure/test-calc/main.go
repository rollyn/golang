package main

import (
  "fmt"
  "github.com/rollyn/calc"
  "strconv"
)

func main() {
  t := strconv.Itoa( calc.Add(5,5) )
  fmt.Printf( t  )
}