package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func HandleError() {
	if r := recover(); r != nil {
		color.Set(color.FgRed)
		fmt.Println(r)
		color.Unset()
	}
}
