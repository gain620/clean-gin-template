package logo

import (
	"clean-gin-template/pkg"
	"fmt"
	"github.com/common-nighthawk/go-figure"
)

func PrintLogo() {
	dev := figure.NewColorFigure("Gain Chang", "", "green", true)
	dev.Print()
	program := figure.NewColorFigure("Gin API Server", "", "blue", true)
	program.Print()
	version := figure.NewColorFigure(fmt.Sprintf("v%s", pkg.GetVersion()), "", "red", true)
	version.Print()
	fmt.Println()
}
