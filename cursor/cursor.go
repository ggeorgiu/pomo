package cursor

import "fmt"

const (
	codeHideCursor      = "\033[?25l"
	codeShowCursor      = "\033[?25h"
	codeMoveToLineStart = "\r\033[K"
)

func Show() {
	fmt.Print(codeShowCursor)
}

func Hide() {
	fmt.Print(codeHideCursor)
}

func MoveToLineStart() {
	fmt.Print(codeMoveToLineStart)
}
