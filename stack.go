package stack

import (
	"bytes"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

// Regular expressions for MinifiedStack().
var (
	fnlineRegexp   = regexp.MustCompile(`^\s*(.+)\([^\)]+\)$`)
	filelineRegexp = regexp.MustCompile(`^\s*(.+:\d+).+$`)
)

// Minify returns a condensed version of debug.Stack().
func Minify(stack []byte) []byte {
	var w bytes.Buffer
	input := strings.Split(strings.TrimSpace(string(stack)), "\n")

	// Copy "goroutine" line.
	if strings.HasPrefix(input[0], "goroutine ") {
		fmt.Fprintln(&w, input[0])
		input = input[1:]
	}

	for i := 0; i < len(input)-1; i += 2 {
		fnline, fileline := input[i], input[i+1]

		// Parse function name.
		fn := fnline
		if a := fnlineRegexp.FindStringSubmatch(fnline); a != nil {
			fn = a[1] + "()"
		}

		// Parse base filename.
		file := fileline
		if a := filelineRegexp.FindStringSubmatch(fileline); a != nil {
			file = filepath.Base(a[1])
		}

		fmt.Fprintf(&w, "%s: %s\n", fn, file)
	}

	return w.Bytes()
}
