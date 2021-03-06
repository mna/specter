package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/mna/specter/vm"
)

func main() {
	var cpuprof string

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\tcmd [-cpu=\"/path/to/cpu.prof\"] [-mem=\"/path/to/mem.prof\"] FILE\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.StringVar(&cpuprof, "cpu", "", "activates the cpu profiler, and saves the data in this file")
	flag.Parse()

	if flag.NArg() > 0 {
		if len(cpuprof) > 0 {
			f, err := os.Create(cpuprof)
			if err != nil {
				panic(err)
			}
			pprof.StartCPUProfile(f)
			defer pprof.StopCPUProfile()
		}
		vm := vm.New()
		if f, err := os.Open(flag.Arg(0)); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		} else {
			vm.Run(f)
		}
	} else {
		fmt.Println("A file name must be specified.")
		flag.Usage()
	}
}
