package main

import (
	"flag"
	"fmt"
	"os"

	_ "image/jpeg"
	_ "image/png"
)

type options struct {
	input  string
	output string
	loop   int
	delay  int
	quiet  bool
}

var (
	loop  int
	delay int
	quiet bool

	o = &options{}
)

func init() {
	flag.StringVar(&o.input, "i", "", "Input source images")
	flag.StringVar(&o.output, "o", "animated.gif", "Output file name")
	flag.IntVar(&loop, "loop", 0, "Loop count (0 == infinite)")
	flag.IntVar(&delay, "delay", 1000, "Delay in milliseconds")
	flag.BoolVar(&quiet, "quiet", false, "Do not output verbose log")
	flag.Parse()
}

func main() {

	if err := run(o); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

}
