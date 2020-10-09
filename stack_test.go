package stack_test

import (
	"bytes"
	"testing"

	"github.com/benbjohnson/stack"
)

func TestMinify(t *testing.T) {
	input := `
goroutine 61 [running]:
runtime/debug.Stack(0xd, 0xd, 0x4f44328)
	/usr/local/go/src/runtime/debug/stack.go:24 +0x9d
github.com/benbjohnson/stack/v2/foo.(*Bar).baz(0xc000c6a5b0, 0x4e00001, 0x0, 0x0, 0x0)
	/src/stack/benbjohnson/stack/foo.go:579 +0x7e
created by net/http.(*Server).Serve
	/usr/local/go/src/net/http/server.go:2962 +0x35c
`[1:]

	if got, want := stack.Minify([]byte(input)), []byte(`
goroutine 61 [running]:
runtime/debug.Stack(): stack.go:24
github.com/benbjohnson/stack/v2/foo.(*Bar).baz(): foo.go:579
net/http.(*Server).Serve(): server.go:2962
`[1:]); !bytes.Equal(got, want) {
		t.Fatalf("Minify()=%q, expected %q", got, want)
	}
}
