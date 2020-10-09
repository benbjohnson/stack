# stack

Go debug/stack utility functions and command line interfaces.


## CLI

### stack-minify

The `stack-minify` CLI command accepts text through STDIN and converts all lines
which have a two-line stack entry format to a single-line minified format.

```sh
$ cat mystack.txt | stack-minify
```


### Utility functions

## Minify()

The `stack.Minify()` function accepts the output of `debug.Stack()` and returns
a compact version that shows one stack entry per line. Function variables and
offsets are removed so it only shows the package, function, filename, & line
number.

```
stack.Minify(debug.Stack())
```

```
Input:

goroutine 61 [running]:
runtime/debug.Stack(0xd, 0xd, 0x4f44328)
	/usr/local/go/src/runtime/debug/stack.go:24 +0x9d
github.com/benbjohnson/stack/v2/foo.(*Bar).baz(0xc000c6a5b0, 0x4e00001, 0x0, 0x0, 0x0)
	/src/stack/benbjohnson/stack/foo.go:579 +0x7e
```

```
Output:

goroutine 61 [running]:
runtime/debug.Stack(): stack.go:24
github.com/benbjohnson/stack/v2/foo.(*Bar).baz(): foo.go:579
```
