package main

import (
  "fmt"
  "os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  file, err := os.Open("ex1.go") // for read access
  check(err)
  defer file.Close()

  // read a few bytes of the file into a slice
  slice1 := make([]byte, 5)
  n1, err := file.Read(slice1)
  check(err)
  fmt.Printf("%d bytes: %s\n", n1, string(slice1))

  // seek to a known location
  o2, err := file.Seek(6, 0)
  check(err)
  slice2 := make([]byte, 2)
  n2, err := file.Read(slice2)
  check(err)
  fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(slice2))
}
