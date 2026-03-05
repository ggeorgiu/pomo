package main

import (
	"fmt"
	"pomo/cursor"
	"pomo/screen"
	"strconv"

	_ "embed"

	"github.com/gen2brain/beeep"
)

//go:embed icons/beep_fitness.png
var icon []byte

// setup holds necessary resources and returns the functionality to release them
func setup() (release func()) {
	cursor.Hide()
	release = func() {
		cursor.Show()
	}

	return release
}

func run(args []string) error {
	defer screen.Wrap()

	var seconds = 10
	if len(args[1]) != 0 {
		val, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("parse args(0), err: %w", err)
		}

		seconds = val

	}
	s, err := screen.New()
	if err != nil {
		return err
	}

	pb, tick := newProgressBar(seconds)
	go pb.run()

	s.Init()
	for range tick {
		s.Update(pb.quant())
	}

	if err := beeep.Alert("Pomo", "Time's Up", icon); err != nil {
		return err
	}

	return nil
}
