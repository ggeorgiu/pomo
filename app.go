package main

import (
	"os"
	"strconv"
	"strings"
	"unicode"

	_ "embed"

	"github.com/gen2brain/beeep"
	"github.com/ggeorgiu/pomo/cursor"
	"github.com/ggeorgiu/pomo/screen"
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

func stop() {
	cursor.MoveToLineStart()
	screen.Print("Interrupted.")
	os.Exit(0)
}

func run(args []string) error {
	defer screen.Wrap()

	var seconds = 10
	if len(args[1]) > 0 {
		seconds = parse(args[1])
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

func parse(duration string) int {
	if unicode.IsDigit(rune(duration[len(duration)-1])) {
		val, err := strconv.Atoi(duration)
		if err != nil {
			panic(err)
		}

		return val
	}

	var multiplier int
	switch {
	case strings.HasSuffix(duration, "s"):
		multiplier = 1
	case strings.HasSuffix(duration, "m"):
		multiplier = 60
	default:
		multiplier = 1
	}

	val, err := strconv.Atoi(duration[:len(duration)-1])
	if err != nil {
		panic(err)
	}

	return val * multiplier
}
