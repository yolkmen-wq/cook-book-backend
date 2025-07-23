package main

import (
	"fmt"
	"os"

	"cook-book-backend/internal/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		fmt.Printf("Failed to create app: %v\n", err)
		os.Exit(1)
	}

	if err := a.Run(); err != nil {
		fmt.Printf("Failed to run app: %v\n", err)
		os.Exit(1)
	}
}