package main

import (
  "testing"
  // "fmt"
)

// func TestCopy(t *testing.T) {
//   CopyFile("./output/" + "index.html", "./templates/index.html", 0755)
// }

func TestCreateSlides(t *testing.T) {
  var fileName = "mocks/test.md"

  slideCount := CreateSlides(fileName)

  if slideCount != 3 {
    t.Error("Expected 3 but got ", slideCount)
  }

}

func TestCreateSite(t *testing.T) {
  CreateSite(3)
}
