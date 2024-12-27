package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"nsight.se/nsight/app/internal/config"
)

var (
	buildTime string
	revision  string
	version   string
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config, err := config.Parse()
	if err != nil {
		panic(fmt.Errorf("error parsing config: %w", err))
	}

	log.Println(map[string]any{
		"buildTime": buildTime,
		"revision":  revision,
		"version":   version,
		"config": map[string]any{
			"addr": config.Address,
			"port": config.Port,
		},
	})

	for {
		select {
		case <-ctx.Done():
			log.Println("exiting")
			return
		case <-time.After(time.Second):
			log.Println("a second has passed")
		}
	}
}
