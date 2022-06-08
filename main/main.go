package main

import (
	"fmt"
	"log"
	"os"
	uic_go_sdl2 "uic-go-sdl2"
	"uic-go-sdl2/gui"
)

func main() {

	if len(os.Args) > 2 { // D:\repo\uic-go-sdl2\gui_raw D:\repo\uic-go-sdl2
		checkIfDir(os.Args[1])
		checkIfDir(os.Args[2])
		fmt.Printf("Compiling UI with paths: \n\t Input Directory: %s \n\t Output Directory: %s \n", os.Args[1], os.Args[2])
		uic_go_sdl2.Compile(os.Args[1], os.Args[2])
		fmt.Println("Done.")
	} else {
		fmt.Println("Expected input and output paths were not given as arguments. Running GUI instead.")
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
