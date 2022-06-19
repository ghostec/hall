package main

import (
	"go.uber.org/zap"
)

func main() {
	log, _ := zap.NewProduction()
	defer log.Sync()

	sv := NewServer()
	addr := "localhost:8080"

	if err := sv.Listen(addr); err != nil {
		log.Fatal("failed to listen", zap.String("addr", addr), zap.Error(err))
	}

	if err := sv.Serve(); err != nil {
		log.Fatal("failed to serve", zap.Error(err))
	}
}
