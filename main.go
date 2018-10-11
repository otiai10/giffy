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
	quiet  bool
)

func init() {
	flag.StringVar(&input, "i", "", "Input source images")
	flag.StringVar(&output, "o", "animated.gif", "Output file name")
	flag.IntVar(&loop, "loop", 0, "Loop count (0 == infinite)")
	flag.IntVar(&delay, "delay", 1000, "Delay in milliseconds")
	flag.BoolVar(&quiet, "quiet", false, "Do not output verbose log")
	flag.Parse()
}

func main() {

	lg := logger{quiet}

	matches, err := filepath.Glob(input)
	if err != nil {
		debug.Println(err)
		return
	}
	if len(matches) == 0 {
		lg.Printf("No files found with `%v`", input)
		return
	}
	lg.Printf("%d files found. Decoding...\n", len(matches))

	g := &gif.GIF{LoopCount: loop}

	for _, name := range matches {
		lg.Printf(name)
		if err := push(g, name); err != nil {
			debug.Println(err)
			return
		}
		lg.Printf("\033[%dD", len(name))
		lg.Printf(strings.Repeat(" ", len(name)))
		lg.Printf("\033[%dD", len(name))
		lg.Printf("✔ ")
	}
	lg.Printf("\n")

	f, err := os.Create(output)
	if err != nil {
		debug.Println(err)
		return
	}

	if err := gif.EncodeAll(f, g); err != nil {
		debug.Println(err)
		return
	}

	lg.Printf("Encoded successfully to %s\n", output)
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

type logger struct {
	quiet bool
}

func (l logger) Printf(format string, v ...interface{}) {
	if quiet {
		return
	}
	fmt.Printf(format, v...)
}
