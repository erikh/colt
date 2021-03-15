package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 2 {
		help()
	}

	switch os.Args[1] {
	case "help", "--help", "-h":
		help()
	}

	args := os.Args[1:]

	var (
		background string
		attrs      []string
	)

	foreground := args[0]

	if len(args) > 1 {
		for i := 0; i < len(args); i++ {
			switch args[i] {
			case "on":
				if len(args[i:]) < 2 {
					err("invalid syntax; 'on' requires a background color")
				}

				i++
				background = args[i]
			case "with":
				if len(args[i:]) < 2 {
					err("invalid syntax; 'with' requires attributes")
				}

				attrs = args[i+1:]
				i = len(args)
			default:
				background = args[i]
			}
		}
	}

	mkColor(foreground, background, attrs)
}

func mkColor(foreground, background string, attrs []string) {
	colorAttrs := []color.Attribute{}

	colorAttrs = append(colorAttrs, bgColorMap[background])
	colorAttrs = append(colorAttrs, colorMap[foreground])

	for _, attr := range attrs {
		colorAttrs = append(colorAttrs, attrMap[attr])
	}

	c := color.New(colorAttrs...)
	c.EnableColor()
	c.Print()
}

func help() {
	fmt.Fprintln(os.Stderr, "https://github.com/erikh/colt/blob/main/README.md")
	os.Exit(1)
}

func err(s string) {
	fmt.Fprintln(os.Stderr, s)
	help()
}
