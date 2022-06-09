package main

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/armnerd/go-skeleton/cmd"
)

var (
	SetCpuCount string
)

// @title go-skeleton
// @version 1.0
// @description Golang 脚手架，Go 简单，Go 直接
// @termsOfService https://github.com/armnerd/go-skeleton
func main() {
	if SetCpuCount != "" {
		procsNum, err := strconv.Atoi(SetCpuCount)
		if err == nil {
			runtime.GOMAXPROCS(procsNum)
			fmt.Printf("GOMAXPROCS num set %v\n", procsNum)
		}
	}
	cmd.Execute()
}
