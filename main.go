package main

import (
	"github.com/dm0275/weather-cli/cmd"
)

func main() {
	core := cmd.NewCLI()
	core.Execute()
}
