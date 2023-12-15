package main

import (
	"ebc_analyzer/pkg/analyzer"
	"ebc_analyzer/util/ui/message"
	"ebc_analyzer/util/ui/message/header"
	"fmt"
	"strings"

	"github.com/alexflint/go-arg"
)

type ECBBlockData struct {
	Unecrypted string
	Encrypted  string
}

type AppArgs struct {
	Action    string       `arg:"-a,--action,required" help:"ECB Attack Types."`
	BlockData ECBBlockData `arg:"-b,--block-data,required" help:"Block data."`
}

func (e *ECBBlockData) UnmarshalText(text []byte) error {
	parts := strings.Split(string(text), ":")
	if len(parts) != 2 {
		return fmt.Errorf("Hatalı format: -e aaaaaaa:bbbbbb")
	}
	e.Unecrypted = parts[0]
	e.Encrypted = parts[1]
	return nil
}

var (
	args AppArgs
)

func main() {

	header.DisplayFigure()

	//7z%2BGu21W2Yi91LAZ1eB8zQ%3D%3D
	//9KG7Vr4LWlpOzvGerl%2BhDbcp%2BiE3K5pw
	//9KG7Vr4LWlr0obtWvgtaWriOhN4BFFJD8Y082QkoxF4%3D

	arg.MustParse(&args)

	switch args.Action {
	case "crack":

		message.Println(message.Warning, 0, "Selected action : crack\n")

		_, err := analyzer.ProcessCracking(args.BlockData.Unecrypted, args.BlockData.Encrypted)
		if err != nil {
			message.Println(message.Error, 0, err.Error())
		}

	case "key_decrypt":
		fmt.Println("Selected action: key_decrypt")
	default:
		fmt.Println("Invalid action:", args.Action)
	}

	return

}
