package main

import (
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"
	"path/filepath"
	"strings"
)

func run(opt *options) error {

	lg := logger{opt.quiet}

	matches, err := filepath.Glob(opt.input)
	if err != nil {
		return err
	}
	if len(matches) == 0 {
		return fmt.Errorf("No files found with `%v`", opt.input)
	}

	lg.Printf("%d files found. Decoding...\n", len(matches))

	g := &gif.GIF{LoopCount: loop}

	for _, name := range matches {
		lg.Printf(name)
		if err := push(g, name); err != nil {
			return err
		}
		lg.Printf("\033[%dD", len(name))
		lg.Printf(strings.Repeat(" ", len(name)))
		lg.Printf("\033[%dD", len(name))
		lg.Printf("✔ ")
	}
	lg.Printf("\n")

	f, err := os.Create(opt.output)
	if err != nil {
		return err
	}

	if err := gif.EncodeAll(f, g); err != nil {
		return err
	}

	lg.Printf("Encoded successfully to %s\n", opt.output)

	return nil
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
