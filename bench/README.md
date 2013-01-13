# Benchmarks

This folder contains commands to run benchmarks comparing the Go implementation (specter) to the original C implementation (TinyVM), and to run the profiler on the Go code.

Make sure you adjust the paths in the Makefile and the shell scripts to point to your installation of each program.

On Mac OSX, the gnu-time (gtime) command is required (`brew install gnu-time`).

## Running

You can run the scripts with this command:

    make bench FILE=fib LOOPS=1000

By default (if neither FILE nor LOOPS are specified), it runs with `fib` and `1000`. The FILE argument is the base example file name to run, without the `.vm` extension. The LOOPS argument is the number of iteration. It uses `time` to capture the execution time.

You can also run all benchmarks (that is, run the benchmark for all .vm files in the ./cmd/examples/ directory) with `make all`.

## Performance

At the moment, Go's implementation is between 1.3 and 1.8 times the C implementation (except `loop.vm` which is much slower, and `nop.vm`, almost 2 times) on a 2012 MacBook Pro Retina (Core i7 2.3GHz, 8Gb RAM). I'm positive optimizations can be made, as this is pretty much a first draft of porting TinyVM's code.

See the `./results/` directory for all the numbers.
