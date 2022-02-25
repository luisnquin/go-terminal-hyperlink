// Package dedicated to creating hyperlinks for CLI programs. Includes colors, 
// and compatibility check. Report any issues.

package link

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	shl "github.com/luisnquin/go-supports-hyperlinks"
)

// Default option
var (
	Stdout = Hyperlink{Stream: os.Stdout}
	Stderr = Hyperlink{Stream: os.Stderr}
)

// Per default the stream uses stdout, less syntax, of course.
// Avoid using it, use New(... *os.File) function
type Hyperlink struct {
	Colors []color.Attribute
	Stream *os.File `default:"os.Stdout"`
}

// Like fmt family but evaluates and transforms the required inputs to a hyperlink
type Hyperlinker interface {
	Sprint(url, text string) string
	Sprintln(url, text string) string
	Sprintf(url, text string, a ...interface{}) string
	Print(url, text string)
	Println(url, text string)
	Printf(url, text string, a ...interface{})
}

// Create a new instance of Hyperlink without create a new variable, just receives an *os.File
func New(stream *os.File) *Hyperlink {
	return &Hyperlink{Stream: stream}
}

// Set any number of color attributes, it returns a Hyperlinker that allows you to have a family of fmt methods
func (h *Hyperlink) SetColors(colors ...color.Attribute) Hyperlinker {
	h.Colors = colors
	return h
}

// Returns a clickable string, works like HTML <a> tags. All of this, without alt attribute, obviusly
func (h *Hyperlink) Sprint(url, text string) string {
	supportsHyperlinks := shl.SupportsHyperlinks(h.Stream)

	switch {
	case len(h.Colors) > 0 && supportsHyperlinks:
		return color.New(h.Colors...).Sprint("\u001B]8;;" + url + "\u0007" + text + "\u001B]8;;\u0007")
	case len(h.Colors) > 0:
		return color.New(h.Colors...).Sprint(text + " ( " + url + " )")
	case supportsHyperlinks:
		return "\u001B]8;;" + url + "\u0007" + text + "\u001B]8;;\u0007"
	default:
		return text + " ( " + url + " )"
	}
}

func (h *Hyperlink) Sprintln(url, text string) string {
	supportsHyperlinks := shl.SupportsHyperlinks(h.Stream)

	switch {
	case len(h.Colors) > 0 && supportsHyperlinks:
		return color.New(h.Colors...).Sprintln("\u001B]8;;" + url + "\u0007" + text + "\u001B]8;;\u0007")
	case len(h.Colors) > 0:
		return color.New(h.Colors...).Sprintln(text + " ( " + url + " )")
	case supportsHyperlinks:
		return "\u001B]8;;" + url + "\u0007" + text + "\u001B]8;;\u0007" + "\n"
	default:
		return text + " ( " + url + " )\n"
	}
}

// Returns the same clickable string, but also allows you to format the text, amazing or not?
func (h *Hyperlink) Sprintf(url, text string, a ...interface{}) string {
	supportsHyperlinks := shl.SupportsHyperlinks(h.Stream)

	switch {
	case len(h.Colors) > 0 && supportsHyperlinks:
		return color.New(h.Colors...).Sprint("\u001B]8;;" + url + "\u0007" + fmt.Sprintf(text, a...) + "\u001B]8;;\u0007")
	case len(h.Colors) > 0:
		return color.New(h.Colors...).Sprintf(fmt.Sprintf(text, a...) + " (" + url + ")")
	case supportsHyperlinks:
		return fmt.Sprintf("\u001B]8;;"+url+"\u0007"+text+"\u001B]8;;\u0007", a...)
	default:
		return fmt.Sprintf(text, a...) + " (" + url + ")"
	}
}

func (h *Hyperlink) Print(url, text string) {
	supportsHyperlinks := shl.SupportsHyperlinks(h.Stream)

	switch {
	case len(h.Colors) > 0 && supportsHyperlinks:
		color.New(h.Colors...).Print("\u001B]8;;" + url + "\u0007" + text + "\u001B]8;;\u0007")
	case len(h.Colors) > 0:
		color.New(h.Colors...).Print(text + " (" + url + ")")
	case supportsHyperlinks:
		fmt.Print("\u001B]8;;" + url + "\u0007" + text + "\u001B]8;;\u0007")
	default:
		fmt.Print(text + " (" + url + ")")
	}
}

func (h *Hyperlink) Println(url, text string) {
	supportsHyperlinks := shl.SupportsHyperlinks(h.Stream)

	switch {
	case len(h.Colors) > 0 && supportsHyperlinks:
		color.New(h.Colors...).Println("\u001B]8;;" + url + "\u0007" + text + "\u001B]8;;\u0007")
	case len(h.Colors) > 0:
		color.New(h.Colors...).Println(text + " (" + url + ")")
	case supportsHyperlinks:
		fmt.Println("\u001B]8;;" + url + "\u0007" + text + "\u001B]8;;\u0007")
	default:
		fmt.Println(text + " (" + url + ")")
	}
}

func (h *Hyperlink) Printf(url, text string, a ...interface{}) {
	supportsHyperlinks := shl.SupportsHyperlinks(h.Stream)

	switch {
	case len(h.Colors) > 0 && supportsHyperlinks:
		color.New(h.Colors...).Printf("\u001B]8;;"+url+"\u0007"+text+"\u001B]8;;\u0007", a...)
	case len(h.Colors) > 0:
		color.New(h.Colors...).Printf(text+" ("+url+")", a...)
	case supportsHyperlinks:
		fmt.Printf("\u001B]8;;"+url+"\u0007"+text+"\u001B]8;;\u0007", a...)
	default:
		fmt.Printf(text+" ("+url+")", a...)
	}
}
