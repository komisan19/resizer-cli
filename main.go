package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
  u, err := uuid.NewRandom()
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(u.String())
}
