// This is a copy and paste (almost everything) from github.com/snugfox/ansi-escapes
// precisely the link function.
// All this without downloading unnecessary things and more, smart, right?

// Have problems with Git Bash(win10), due to the default configuration of that tool.
// Anyway, report any issues.

package link

import "fmt"

// Returns a clickable string, works like HTML <a> tags. All of this, without alt attribute, obviusly
func Sprint(url, text string) string {
	return "\u001B]8;;" + url + "\u0007" + text + "\u001B]8;;\u0007"
}

// Returns the same clickable string, but also allows you to format the text, amazing or not?
func Sprintf(url, text string, a ...interface{}) string {
	return fmt.Sprintf("\u001B]8;;"+url+"\u0007"+text+"\u001B]8;;\u0007", a...)
}
