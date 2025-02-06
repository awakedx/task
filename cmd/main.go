package main

import (
	"log"

	"github.com/awakedx/task/internal/app"
)

func main() {
	if err := app.StartServer(); err != nil {
		log.Fatal(err)
	}
}
