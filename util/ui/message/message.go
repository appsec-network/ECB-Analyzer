package message

import (
	"fmt"

	"github.com/gookit/color"
)

type MessageType int

const (
	Info MessageType = iota
	Warning
	Error
	Success
	Notify
	NoType
)

func Println(messageType MessageType, tabs int, message string, args ...interface{}) {
	var prefix string
	var s color.Style

	switch messageType {
	case Info:
		prefix = "[i]"
		s = color.New(color.Bold, color.FgBlue)
	case Warning:
		prefix = "[!]"
		s = color.New(color.Bold, color.FgYellow)
	case Error:
		prefix = "[-]"
		s = color.New(color.Bold, color.FgRed)
	case Success:
		prefix = "[+]"
		s = color.New(color.Bold, color.FgGreen)
	case Notify:
		prefix = "[*]"
		s = color.New()
	case NoType:
		prefix = ""
		s = color.New()
	default:
		prefix = "[*]"
		s = color.New()
	}

	// Add tabs
	tabsStr := ""
	for i := 0; i < tabs; i++ {
		tabsStr += "    "
	}

	// Add prefix to the message
	if !(prefix == "" || messageType == NoType) {
		message = " " + message
	}

	if messageType == Error {
		message += "\n"
	}

	// This will print styled text to the terminal
	s.Printf(fmt.Sprintf("%s%s%s", tabsStr, prefix, message), args...)
}
