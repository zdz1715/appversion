package main

import (
	"fmt"

	goappversion "github.com/zdz1715/go-app-version"
)

func main() {
	goappversion.SetVersion("v0.1.1")

	fmt.Printf("%#v", goappversion.Get())
}
