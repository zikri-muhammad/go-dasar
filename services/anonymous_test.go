package services

import (
	"fmt"
	"testing"
)

type Backlists func(string) bool

func register(name string, backlist Backlists) {
	nameFiltered := backlist(name)
	if nameFiltered {
		fmt.Println("You are backlisted", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func TestRegistrasiUsers(t *testing.T) {
	backlist := func(name string) bool {
		return name == "Anjing"
	}

	register("Anjing", backlist)
	register("Muhammad", backlist)
}
