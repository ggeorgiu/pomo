package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ggeorgiu/pomo/cursor"
	"github.com/ggeorgiu/pomo/screen"
)

func main() {
	release := setup()
	defer release()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		screen.Print("Interrupted.")
		cursor.MoveToLineStart()
		release()
	
		os.Exit(0)
	}()

	if err := run(os.Args); err != nil {
		slog.Error("err", "err", err)
	}
}
