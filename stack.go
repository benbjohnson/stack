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
	filelineRegexp = regexp.MustCompile(`^\s*(.+:\d+).*$`)
)

// Minify returns a condensed version of debug.Stack().
func Minify(stack []byte) []byte {
	var w bytes.Buffer
	input := strings.Split(strings.TrimSpace(string(stack)), "\n")

	for len(input) > 0 {
		// Read next two lines and ensure they match our regexes.
		var fnsubmatch []string
		if strings.HasPrefix(input[0], "created by ") {
			fnsubmatch = []string{"", strings.TrimPrefix(input[0], "created by ")}
		} else {
			fnsubmatch = fnlineRegexp.FindStringSubmatch(input[0])
		}

		var filesubmatch []string
		if len(input) >= 2 {
			filesubmatch = filelineRegexp.FindStringSubmatch(input[1])
		}

		// If the two line pair doesn't match or if we only have one line,
		// then simply print the current line and move ahead.
		if len(fnsubmatch) == 0 || len(filesubmatch) == 0 {
			fmt.Fprintln(&w, input[0])
			input = input[1:]
			continue
		}

		// Print single line on match and skip the next two input lines.
		fmt.Fprintf(&w, "%s(): %s\n", fnsubmatch[1], filepath.Base(filesubmatch[1]))
		input = input[2:]
	}

	return w.Bytes()
}
