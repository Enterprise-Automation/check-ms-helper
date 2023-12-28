package checks

import (
	"fmt"
	"os"
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
	for _, check := range checks {
		if check.name == os.Getenv("ACTION") {
			fmt.Println(check.function())
		}
	}
}
