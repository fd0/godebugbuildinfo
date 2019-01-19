package main

import (
	"fmt"
	"runtime/debug"

	"github.com/fatih/color"
)

func main() {
	color.Green("test program for debug.BuildInfo")

	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("not ok")
	} else {
		fmt.Printf("path: %v\n  main: %#v\n  deps:\n", info.Path, info.Main)
		for _, dep := range info.Deps {
			fmt.Printf("    %#v\n", dep)
		}
	}

}
