package main

import link "github.com/luisnquin/go-terminal-link"

func main() {
	url := link.Sprint("https://github.com/luisnquin/go-terminal-link", "Go terminal link")
	println(url)
}
