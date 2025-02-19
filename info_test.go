package appversion

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	t.Logf("%#v", Get())
	SetVersion("v1.22.3")
	t.Logf("%#v", Get())
	SetVersion("v1.23.2")
	SetName("Kubernetes")
	fmt.Println(Get().Json())
}
