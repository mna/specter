# Specter

Specter is an implementation of [GenTiradentes' TinyVM][tvm] in Go (the original is written in C).

This is very much to learn about virtual machines (this is my first attempt at a VM, so huge thanks to GenTiradentes for making a minimal one, easy to grasp).

It runs all examples available in TinyVM's repository, at [roughly 30% slower than C][bench], and at 7% slower with bounds-checking disabled.

## Thanks

The great Go community on the [mailing list][nuts].

## License

The [BSD 3-Clause license][bsd].

[bsd]: http://opensource.org/licenses/BSD-3-Clause
[tvm]: https://github.com/GenTiradentes/tinyvm
[bench]: https://github.com/PuerkitoBio/specter/tree/master/bench
[nuts]: https://groups.google.com/forum/#!topic/golang-nuts/XhK5tGUsZnQ
