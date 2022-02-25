## Go-terminal-hyperlink

Create clickable hyperlinks in the terminal

![Result](https://media.giphy.com/media/soyIDSLKcTgG6k7tcr/giphy.gif)

## Install

```console
$ go get -u github.com/luisnquin/go-terminal-hyperlink
```

## Usage

```go
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

```

## Why?

- Some terminals are not fully compatible, sometimes with versions impossible to infer despite the existence of future compatible ones. [Check this](https://github.com/jamestalmage/supports-hyperlinks/pull/8) for more

[![dasdqanoduqbwd.png](https://i.postimg.cc/L8trKF2X/dasdqanoduqbwd.png)](https://postimg.cc/DSzCLDWk)



 - Sometimes the incompatibility could come from the mere fact that the terminal interprets it as invisible characters, offering results like:

[![issue-2.png](https://i.postimg.cc/BZC4rzcM/issue-2.png)](https://postimg.cc/3WWspnqv)

To prevent that issues, on incompatible terminals, the result is replaced by:

[![but-1.png](https://i.postimg.cc/j517s9tw/but-1.png)](https://postimg.cc/1nptBvgs)

Obviusly, this without force the behavior as the first image of the section.

Still, if you just want to force this without setting a new flag, put it in your code before
to use the package
```go
_ = flag.Set("hyperlinks", "true")
```

## API
It's about assimilating to projects created in JavaScript like:
 - https://github.com/sindresorhus/terminal-link
 - https://github.com/jamestalmage/supports-hyperlinks

So this [guide(supported terminals)](https://gist.github.com/egmontkob/eb114294efbcd5adb1944c9f3cb5feda) might not be excluded

### Variables

- Stdout

```go
package main

import link "github.com/luisnquin/go-terminal-hyperlink"

func main() {
    moreText := "Yes, more text"
    println(link.Stdout.Sprintf(url, text, moreText))
}
```

- Stderr

```go
package main

import link "github.com/luisnquin/go-terminal-hyperlink"

func main() {
    link.Stderr.Println(url, text)
}
```

### Functions

- New

```go
package main

import (
    "os"

    link "github.com/luisnquin/go-terminal-hyperlink"
)

func main() {
    println(link.New(os.Stdout).Sprintln(url, text))
}
```

### Methods
 - SetColors
```go
package main

import (
    "os"

    "github.com/fatih/color"
    link "github.com/luisnquin/go-terminal-hyperlink"
)

func main() {
    link.Stdout.SetColors(color.FgHiGreen, color.ReverseVideo).Println(url, text)
}
```


### Structs

- Hyperlink

  #### Fields

  - Colors `[]color.Attribute`
  - Stream `*os.File`

  #### Methods
  - SetColors(colors ...color.Attribute) Hyperlinker
  - Sprint(url, text string) string
  - Sprintln(url, text string) string
  - Sprintf(url, text string, a ...interface{}) string
  - Print(url, text string)
  - Println(url, text string)
  - Printf(url, text string, a ...interface{})

  **Note**: By default Stream is os.Stdout, this field is required if you want to create a new instance

### Interfaces

- Hyperlinker
  - Sprint(url, text string) string
  - Sprintln(url, text string) string
  - Sprintf(url, text string, a ...interface{}) string
  - Print(url, text string)
  - Println(url, text string)
  - Printf(url, text string, a ...interface{})

## A cat
![Result](https://media.giphy.com/media/vFKqnCdLPNOKc/giphy.gif)

## License
MIT © [Luis Quiñones Requelme](https://github.com/luisnquin)