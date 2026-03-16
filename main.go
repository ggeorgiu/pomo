package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	release := setup()
	defer release()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c

		release()
		stop()
	}()

	if err := run(os.Args); err != nil {
		slog.Error("err", "err", err)
	}
}
