package config

import (
	"os"
	"strings"
)

var AppRoot string

func SetAppRoot(cmd string) {
	var goRunMain = false
	for _, v := range [4]string{"var", "tmp", "Temp", "./main"} {
		if strings.Contains(cmd, v) {
			goRunMain = true
			break
		}
	}
	if goRunMain {
		AppRoot, _ = os.Getwd()
	} else {
		AppRoot = strings.Replace(cmd, "/main", "", -1)
	}
}
