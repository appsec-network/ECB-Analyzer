package header

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

func DisplayFigure() {
	figure.NewFigure("EBC ANALYZER", "cybermedium", true).Print()
	fmt.Println()
}
