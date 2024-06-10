package main

import (
	"log/slog"
	"os"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	logger.Info("A group of walrus emerges from the ocean",
		slog.Int("size", 10),
		slog.String("animal", "walrus"),
	)
}
