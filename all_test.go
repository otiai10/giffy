package main

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestRun(t *testing.T) {

	opt := &options{
		input:  "./test/data/case00/*.png",
		output: "./test/out/animated00.gif",
	}
	err := run(opt)
	Expect(t, err).ToBe(nil)

	When(t, "input not specified", func(t *testing.T) {
		opt := &options{}
		err := run(opt)
		Expect(t, err).Not().ToBe(nil)
	})

}
