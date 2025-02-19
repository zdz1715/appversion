package main

import (
	"fmt"

	goappversion "github.com/zdz1715/go-app-version"
)

func main() {
	fmt.Printf("%#v", goappversion.Get())
}
