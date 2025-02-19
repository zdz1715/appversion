package main

import (
	"fmt"

	"github.com/zdz1715/appversion"
)

func main() {
	appversion.SetVersion("v0.1.1")

	fmt.Println(appversion.Get().Json())
}
