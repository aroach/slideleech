package main

import (
  "testing"
  "fmt"
)

func TestCopy(t *testing.T) {
  fmt.Println("test")
  CopyFile("./output/" + "index.html", "./templates/index.html", 0755)
}
