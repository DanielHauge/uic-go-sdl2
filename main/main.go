package main

import (
	"fmt"
	"log"
	"os"
	uic_go_sdl2 "uic-go-sdl2"
	gui "uic-go-sdl2/TestOutput"
)

func main() {

	if len(os.Args) > 2 { // D:\\repo\\uic-go-sdl2\\TestInput D:\\repo\\uic-go-sdl2\\TestOutput
		fmt.Println("Compiling UI with paths")
		checkIfDir(os.Args[1])
		checkIfDir(os.Args[2])
		fmt.Println("Input Directory: ", os.Args[1])
		fmt.Println("Output Directory: ", os.Args[2])
		uic_go_sdl2.Compile(os.Args[1], os.Args[2])
		fmt.Println("Done.")
	} else {
		fmt.Println("Expected input and output paths as argument.")
		gui.RunGui()
	}
}

func checkIfDir(dir string) {
	fInfo, err := os.Stat(dir)
	if err != nil {
		panic(err)
	}
	if !fInfo.IsDir() {
		log.Panicf("%s was not a directory", dir)
	}
}
