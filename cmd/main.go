package main

import (
	"fmt"
	"github.com/PuerkitoBio/specter/vm"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		vm := vm.New()
		if f, err := os.Open(os.Args[1]); err != nil {
			panic(err)
		} else {
			vm.Run(f)
		}
	} else {
		fmt.Println("A file name must be specified.")
		fmt.Println("Usage: specter FILE")
	}
}
