// Package main is the command center of the app
package main

import (
	"os"

	"Go-Do/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
