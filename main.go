package checks

import (
	"fmt"
)

type Function func() interface{}

type check struct {
	name     string
	function Function
}

var checks []check

func NewCheck(name string, function Function) {
	fmt.Println("New Check", name)
	checks = append(checks, check{name: name, function: function})
}

func RegisterChecks() {
	// thingToCheck := "check2"
	// call the right function
	fmt.Println(checks)
}
