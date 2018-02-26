package main

import (
	"os"
	"path"
	"base64"
)


func main() {
	if len(os.Args) != 2 {
		println("Usage: " + path.Base(os.Args[0]) + " <input>")
		os.Exit(1)
	}
	
	b64 := base64.Encode(&os.Args[1])
	println(*b64)
}