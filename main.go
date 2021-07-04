package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

type Extension struct {
	extension string
}

func (e *Extension) CreatePath()string {
	uu, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
	}
  newPath := uu.String() + "." + e.extension
	return newPath
}

func main() {
	flag.Parse()
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

  img, data, err := image.Decode(file)
  if err != nil {
      log.Fatal(err)
  }
  m := resize.Resize(100, 0, img, resize.Lanczos3)
  switch data {
  case "png":
    extension := &Extension{"png"}
    out, err := os.Create(extension.CreatePath())
    if err != nil {
      log.Fatal(err)
    }
    defer out.Close()
    png.Encode(out, m)
  case "jpeg":
    extension := &Extension{"jpeg"}
    out, err := os.Create(extension.CreatePath())
    if err != nil {
      log.Fatal(err)
    }
    defer out.Close()
    jpeg.Encode(out, m, nil)
  case "gif":
    extension := &Extension{"gif"}
    out, err := os.Create(extension.CreatePath())
    if err != nil {
      log.Fatal(err)
    }
    defer out.Close()
    gif.Encode(out, m, nil)
  default:
    log.Fatal("sorry, it's not support")
  }
}
