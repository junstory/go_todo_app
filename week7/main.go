package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/junstory/go_todo_app/week6/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}

// 리팩토링한 run 함수
func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)
	mux, cleanup, err := NewMux(ctx, cfg)
	defer cleanup()
	if err != nil {

	}
	s := NewServer(l, mux)
	return s.Run(ctx)
}
