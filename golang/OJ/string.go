package main

import (
	"fmt"
	"strings"
)

func stringFormatBoth(fg, bg int, str string, args []string) string {
	return fmt.Sprintf("\x1b[48;5;%dm\x1b[38;5;%d;%sm%s\x1b[0m", bg, fg, strings.Join(args, ";"), str)
}

func frameText(text string) string {
	return stringFormatBoth(15, 0, text, []string{"1"})
}
