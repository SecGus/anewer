package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var quietMode bool
	var dryRun bool
	var trim bool
	var output string
	flag.BoolVar(&quietMode, "q", false, "quiet mode (no output at all)")
	flag.BoolVar(&dryRun, "d", false, "don't append anything to the file, just print the new lines to stdout")
	flag.BoolVar(&trim, "t", false, "trim leading and trailing whitespace before comparison")
	flag.StringVar(&output, "o", "", "define a file to write the new output to")
	flag.Parse()

	fns := flag.Args()

	lines := make(map[string]bool)

	var f io.WriteCloser
	var err error

	for _, fn := range fns {
		// read the whole file into a map if it exists
		r, err := os.Open(fn)
		if err == nil {
			sc := bufio.NewScanner(r)

			for sc.Scan() {
				if trim {
					lines[strings.TrimSpace(sc.Text())] = true
				} else {
					lines[sc.Text()] = true
				}
			}
			r.Close()
		}
	}

	// Check if -o option was provided
	if output != "" {
		// open the output file for appending new stuff
		f, err = os.OpenFile(output, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open output file for writing: %s\n", err)
			return
		}
		defer f.Close()
	} else if !dryRun && len(fns) > 0 {
		// re-open the last file for appending new stuff
		f, err = os.OpenFile(fns[len(fns)-1], os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open file for writing: %s\n", err)
			return
		}
		defer f.Close()
	}

	// read the lines, append and output them if they're new
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		line := sc.Text()
		if trim {
			line = strings.TrimSpace(line)
		}
		if lines[line] {
			continue
		}

		// add the line to the map so we don't get any duplicates from stdin
		lines[line] = true

		if !quietMode {
			fmt.Println(line)
		}
		if !dryRun && f != nil {
			fmt.Fprintf(f, "%s\n", line)
		}
	}
}
