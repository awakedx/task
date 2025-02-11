package main

import (
	"log/slog"

	"github.com/awakedx/task/internal/app"
)

func main() {
	if err := app.StartServer(); err != nil {
		slog.Error("Eror Starting server,", "err:", err.Error())
	}
}
