package main

import (
	"ebc_analyzer/pkg/analyzer"
	"ebc_analyzer/util/ui/message"
	"ebc_analyzer/util/ui/message/header"
	"fmt"
	"strings"

	"github.com/alexflint/go-arg"
)

type BlockData struct {
	Unecrypted string
	Encrypted  string
}

type BlockSize string

type AppArgs struct {
	Action    string    `arg:"-a,--action,required" help:"ECB Attack Types."`
	BlockData BlockData `arg:"-d,--block-data,required" help:"Block data."`
	BlockSize BlockSize `arg:"-s,--block-size" default:"auto" help:"Block Size: 8, 16, 32, 64 or auto."`
	Verbose   bool      `arg:"-v,--verbose" help:"verbosity level"`
}

func (d *BlockData) UnmarshalText(text []byte) error {

	parts := strings.Split(string(text), ":")
	if len(parts) != 2 {
		return fmt.Errorf("Incorrect format: -d sample:c2FtcGxl")
	}
	d.Unecrypted = parts[0]
	d.Encrypted = parts[1]

	return nil
}

func (bs *BlockSize) UnmarshalText(text []byte) error {

	validSizes := []string{"8", "16", "32", "64", "auto"}
	inputSize := string(text)

	for _, size := range validSizes {
		if size == inputSize {
			*bs = BlockSize(inputSize)
			return nil
		}
	}

	return fmt.Errorf("Invalid block size: %s", inputSize)
}

var (
	args AppArgs
)

func main() {

	header.DisplayFigure()

	arg.MustParse(&args)

	switch args.Action {
	case "crack":

		message.Println(message.Warning, 0, "Selected action : crack\n")

		_, err := analyzer.ProcessCracking(args.BlockData.Unecrypted, args.BlockData.Encrypted, string(args.BlockSize))
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
