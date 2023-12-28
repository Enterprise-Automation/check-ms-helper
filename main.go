package checks

import (
	"fmt"
)

type Function func() interface{}

func NewCheck(name string, function Function) {
	fmt.Println("New Check", name)
	function()
}

func RegisterChecks() {
	fmt.Println("Register")
}
