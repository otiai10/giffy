package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"os"
	"path/filepath"
	"strings"

	"github.com/otiai10/debug"

	"image/color/palette"
	"image/gif"
	_ "image/jpeg"
	_ "image/png"
)

var (
	input  string
	output string
	loop   int
	delay  int
)

func init() {
	flag.StringVar(&input, "i", "", "Input source images")
	flag.StringVar(&output, "o", "animated.gif", "Output file name")
	flag.IntVar(&loop, "loop", 0, "Loop count (0 == infinite)")
	flag.IntVar(&delay, "delay", 1000, "Delay in milliseconds")
	flag.Parse()
}

func main() {

	matches, err := filepath.Glob(input)
	if err != nil {
		debug.Println(err)
		return
	}
	if len(matches) == 0 {
		fmt.Printf("No files found with `%v`", input)
		return
	}
	fmt.Printf("%d files found. Decoding...\n", len(matches))

	g := &gif.GIF{LoopCount: loop}

	for _, name := range matches {
		fmt.Print(name)
		if err := push(g, name); err != nil {
			debug.Println(err)
			return
		}
		fmt.Printf("\033[%dD", len(name))
		fmt.Print(strings.Repeat(" ", len(name)))
		fmt.Printf("\033[%dD", len(name))
		fmt.Printf("✔ ")
	}
	fmt.Print("\n")

	f, err := os.Create(output)
	if err != nil {
		debug.Println(err)
		return
	}

	if err := gif.EncodeAll(f, g); err != nil {
		debug.Println(err)
		return
	}

	fmt.Println("Encoded successfully to", output)
}

func push(dest *gif.GIF, filename string) error {

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}
	p := image.NewPaletted(img.Bounds(), palette.Plan9)
	draw.Draw(p, p.Bounds(), img, img.Bounds().Min, draw.Over)
	dest.Image = append(dest.Image, p)
	dest.Delay = append(dest.Delay, delay/10)
	return nil
}
