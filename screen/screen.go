package screen

import (
	"fmt"

	"golang.org/x/term"

	"github.com/ggeorgiu/pomo/color"
)

const bar = "\u275A"

type Screen struct {
	barLen int
	ratio  float32
}

func New() (*Screen, error) {
	w, _, err := term.GetSize(0)
	if err != nil {
		return nil, err
	}

	return &Screen{
		barLen: w / 4,
		ratio:  float32(w) / float32(4) / float32(100),
	}, nil
}

func (s *Screen) Init() {
	fmt.Print("\r")
	fmt.Print("[ ")
	for i := 0; i < s.barLen; i++ {
		fmt.Print(color.Cyan, ".")
	}
	fmt.Print(" ]")
	fmt.Print("[ ", 0, "% ]")
}

func (s *Screen) Update(progress float32) {
	fmt.Print("\r")
	fmt.Print("[ ")
	for i := 0; i < s.barLen; i++ {
		if float32(i) <= s.ratio*progress-1 {
			fmt.Print(color.Green, bar, color.Reset)
			continue
		}
		fmt.Print(color.Cyan, ".", color.Reset)
	}
	fmt.Print(" ]")
	fmt.Print(fmt.Sprintf("[ %.2f%% ]", progress))
}

func Print(text string) {
	fmt.Println(text)
}

func Wrap() {
	fmt.Println()
}
