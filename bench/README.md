# Benchmarks

This folder contains commands to run benchmarks comparing the Go implementation (specter) to the original C implementation (TinyVM).

Make sure you adjust the paths in the shell scripts to point to your installation of each program.

## Running

You can run the scripts with this command:

    make bench FILE=fib LOOPS=1000

By default (if neither FILE nor LOOPS are specified), it runs with `fib` and `1000`. The FILE argument is the base example file name to run, without the `.vm` extension. The LOOPS argument is the number of iteration. It uses `time` to capture the execution time.

## Performance

At the moment, Go's implementation is between 1.3 and 1.8 times the C implementation (except `loop.vm` which is 3 times slower) on a 2012 MacBook Pro Retina (Core i7 2.3GHz, 8Gb RAM). I'm positive optimizations can be made, as this is pretty much a first draft of porting TinyVM's code.
