package link_test

import (
	"fmt"
	"os"
	"testing"

	link "github.com/luisnquin/go-terminal-link"
)

var (
	url  string = "https://www.amazon.com/?&tag=googleglobalp-20"
	text string = "Amazon"
)

func TestUrl(t *testing.T) {
	link := link.Sprint(url, text)

	if _, err := fmt.Fprintln(os.Stdout, link); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
