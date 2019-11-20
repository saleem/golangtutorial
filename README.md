# Installation

[https://ahmadawais.com/install-go-lang-on-macos-with-homebrew/](https://ahmadawais.com/install-go-lang-on-macos-with-homebrew/)

I started with `brew install go` and the Go VSCode plugin.

I had to modify PATH and set GOPATH and GOROOT to make the sublime go plugin to work.

    export GOROOT="$(brew --prefix golang)/libexec"
    export PATH="$PATH:${GOROOT}/bin"
    export GOPATH="${HOME}/.go"

As [this document](https://github.com/golang/go/wiki/GOPATH) says, you only need to set one GOPATH. I created that using this one-time shell command, **after** defining GOPATH as described above:
`test -d "${GOPATH}" || mkdir "${GOPATH}"`

[https://github.com/golang/sublime-build](https://github.com/golang/sublime-build) for Sublime 

After this, I can run simple Go scripts directly in Sublime edit by usin the `⌘B` key-chord.

# Tutorials

[https://gobyexample.com](https://gobyexample.com)

[https://tour.golang.org/basics/](https://tour.golang.org/basics/)

[http://www.golangbootcamp.com/book/](http://www.golangbootcamp.com/book/)

# References
[https://github.com/golang/sublime-build/blob/master/docs/configuration.md](https://github.com/golang/sublime-build/blob/master/docs/configuration.md)

[https://golang.org/cmd/go/#hdr-Compile_and_run_Go_program](https://golang.org/cmd/go/#hdr-Compile_and_run_Go_program)

[https://blog.golang.org/go-slices-usage-and-internals](https://blog.golang.org/go-slices-usage-and-internals)

[https://www.golangprograms.com/](https://www.golangprograms.com/)


