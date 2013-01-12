# Benchmarks

This folder contains commands to run benchmarks comparing the Go implementation (specter) to the original C implementation (TinyVM).

Make sure you adjust the paths in the shell scripts to point to your installation of each program.

## Running

You can run the scripts with this command:

    make bench FILE=fib LOOPS=1000

By default (if neither FILE nor LOOPS are specified), it runs with `fib` and `100`. The FILE argument is the base example file name to run, without the `.vm` extension. The LOOPS argument is the number of iteration. It uses `time` to capture the execution time.

## Performance

At the moment, Go's implementation is a little less than twice the time of the C implementation. I'm positive optimizations can be made, as this is pretty much a first draft of porting TinyVM's code.
