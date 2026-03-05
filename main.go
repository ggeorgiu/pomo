package main

import (
	"log/slog"
	"os"
	"os/signal"
	"pomo/cursor"
	"pomo/screen"
	"syscall"
)

func main() {
	release := setup()
	defer release()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c

		cursor.MoveToLineStart()
		screen.Print("Interrupted.")

		release()
		os.Exit(0)
	}()

	if err := run(os.Args); err != nil {
		slog.Error("err", "err", err)
	}
}
