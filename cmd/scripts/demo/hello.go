package demo

import (
	"fmt"
)

func helloHandler(args []string) error {
	fmt.Println(args)
	fmt.Println("Hello")
	return nil
}
