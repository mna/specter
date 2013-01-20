# Specter

Specter is an implementation of [GenTiradentes' TinyVM][tvm] in Go (the original is written in C).

This is very much to learn about virtual machines (this is my first attempt at a VM, so huge thanks to GenTiradentes for making a minimal one, easy to grasp).

The whole implementation takes about 500 lines of Go code.

## Performance

It runs all examples available in TinyVM's repository, at roughly 30% slower than C, and at 7% slower with bounds-checking disabled. See more about the benchmarks in the [bench subfolder][bench].

## Memory

Specter runs at a higher (but stable) memory footprint than its C counterpart, however note the following (copied [from the mailing list][nuts], thanks to Carlos Castillo, edited):

> BTW: The memory use [...] hovers at a high-ish number (70Mb in my case) because the garbage collection code doesn't run until memory use exceeds a given threshold. When it does, all the generated intermediate strings/slices are found but the memory is not returned to the OS, it is instead re-used for any following allocations, so the memory usage stat according to the OS stays at that threshold. If you set the environment variable GOGCTRACE=1 before running your code you can see when the garbage collector runs, and what happened during the run.

## Thanks

The great Go community on the [mailing list][nuts].

## License

The [BSD 3-Clause license][bsd].

[bsd]: http://opensource.org/licenses/BSD-3-Clause
[tvm]: https://github.com/GenTiradentes/tinyvm
[bench]: https://github.com/PuerkitoBio/specter/tree/master/bench
[nuts]: https://groups.google.com/forum/#!topic/golang-nuts/XhK5tGUsZnQ
