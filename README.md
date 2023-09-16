# Hello, world!

This repo started off as a way to start learning [Golang](https://go.dev/), but when I discovered that it has cross-compilers for several dozen architectures built into it, I got side-tracked with figuring out how to create binaries for other architectures, and then with how to use [`make`](https://www.gnu.org/software/make/) to make it *easy* to create binaries.

So I'm putting this up on Github so people can see the `Makefile`, which shows ...

* How to build a binary, using `make`.

    I know, that's nothing special ... but I've seen a lot of Golang projects out there which don't *have* a `Makefile`, and instead use a home-brewed shell script, or tell the user to manually type a specific "`go build`" command with a bunch of weird options.

    If this helps one Golang programmer figure out how to use `make`, my work here is done.

* How to build binaries for a *list* of architectures, using `make all`.

* How to add "`.exe`" to ms-windows executables' filenames.

* How to embed things like the time it was compiled, a version number maintained outside the source code, and the git hash it was built from, into the binaries when compiling them.
