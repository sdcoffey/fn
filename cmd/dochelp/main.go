package main

import (
	"bufio"
	"fmt"
	"github.com/sdcoffey/fn"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		println("usage: dochelp [fnname|fnname-example]")
		os.Exit(1)
	}
	command := os.Args[1]

	switch command {
	case "fnname":
		FuncNames()
	case "extract-doc":
		ExtractDoc()
	}
}

func forEachStdIn(f func(line string)) {
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		f(reader.Text())
	}
}

// FuncNames takes a declaration returned by go doc, ex:
// func A[T any](abcd string) int -> A
func FuncNames() {
	forEachStdIn(func(line string) {
		if fnname, ok := funcname(line); ok {
			fmt.Fprintln(os.Stdout, fnname)
		}
	})
}

func ExtractDoc() {
	forEachStdIn(func(line string) {
		if strings.HasPrefix(line, " ") {
			fmt.Fprintln(os.Stdout, strings.TrimSpace(line))
		}
	})
}

func funcname(docline string) (string, bool) {
	if strings.HasPrefix(docline, "func ") {
		withoutPrefix := docline[5:]
		parseUntil := fn.Min(
			fn.Reject([]int{strings.Index(withoutPrefix, "["), strings.Index(withoutPrefix, "(")}, func(item int, index int) bool {
				return item < 0
			})...)

		return withoutPrefix[:parseUntil], true
	}
	return "", false
}
