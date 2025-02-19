package main

import (
	"fmt"

	"github.com/zdz1715/appversion"
)

func main() {
	fmt.Println(appversion.Get().Json())
}
