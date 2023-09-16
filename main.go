package main

import (
	"fmt"
	"github.com/dm0275/weather-cli/cmd"
	"os"
)

func main() {
	// TODO
	err := cmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing the CLI '%s'", err)
		os.Exit(1)
	}
}
