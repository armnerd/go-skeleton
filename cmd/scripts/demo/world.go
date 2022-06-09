package demo

import (
	"fmt"
)

func worldHandler(args []string) error {
	fmt.Println(args)
	fmt.Println("world")
	return nil
}
