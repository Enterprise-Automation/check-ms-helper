package checks

import (
	"fmt"
)

func NewCheck(name string) {
	fmt.Println("New Check", name)
}

func RegisterChecks() {
	fmt.Println("Register")
}
