package main

import (
	"VK-Pilot-Project/internal/app"
	"os"
)

func main() {
	if err := app.Run(); err != nil {
		os.Exit(1)
	}
}
