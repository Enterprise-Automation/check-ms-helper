package checks

import (
	"fmt"
)

type function func() interface{}

func NewCheck(name string, function function) {
	fmt.Println("New Check", name)
	function()
}

func RegisterChecks() {
	fmt.Println("Register")
}
