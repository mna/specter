# Benchmarks

This folder contains commands to run benchmarks comparing the Go implementation (specter) to the original C implementation (TinyVM), and to run the profiler on the Go code.

Make sure you adjust the paths in the Makefile and the shell scripts to point to your installation of each program.

On Mac OSX, the gnu-time (gtime) command is required (`brew install gnu-time`).

## Running

You can run the scripts with this command:

    make bench FILE=fib LOOPS=1000

By default (if neither FILE nor LOOPS are specified), it runs with `fib` and `1000`. The FILE argument is the base example file name to run, without the `.vm` extension. The LOOPS argument is the number of iteration. It uses `[g]time` to capture the execution time.

You can also run all benchmarks (that is, run the benchmark for all .vm files in the ./cmd/examples/ directory) with `make all`. To compare results (make sure that the output of specter is the same as the output of TinyVM), run `make cmp`.

## Performance

All tests are run on a 2012 MacBook Pro Retina (Core i7 2.3GHz, 8Gb RAM).

See the `./results/` directory for the raw numbers. I keep a spreadsheet with the changes and the results [here][drive].

You can follow the discussion on the golang-nuts mailing list [here][golang].

[drive]: https://docs.google.com/spreadsheet/ccc?key=0Atx1KnJmATDcdEcweWdGOHdld2lVajlaN0VRbXN6MUE
[golang]: https://groups.google.com/forum/#!topic/golang-nuts/XhK5tGUsZnQ