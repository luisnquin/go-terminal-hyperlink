package main

import (
	"os"

	"github.com/fatih/color"
	link "github.com/luisnquin/go-terminal-hyperlink"
)

var (
	url  = "https://github.com/luisnquin/go-terminal-hyperlink"
	text = "Go terminal hyperlink"
)

func main() {
	// flag.Set("hyperlinks", "true")

	link.Stderr.Println(url, text)

	println(link.Stdout.Sprintf(url, text+" (%s)", "this comes with the text"))

	link.Stdout.SetColors(color.FgHiGreen).Println(url, text)

	link.New(os.Stdout).Println(url, text)

	link.New(os.Stderr).SetColors(color.BgWhite, color.FgHiMagenta).Println(url, text)

	new(link.Hyperlink).Println(url, text) // Default -> os.Stdout
}
