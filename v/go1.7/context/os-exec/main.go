package main

import (
	"context"
	"log"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "sleep", "2")
	if err := cmd.Run(); err != nil {
		log.Fatal(err) // 2016/07/13 16:01:24 signal: killed
	}
}
