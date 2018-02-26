package main

import (
	"fmt"
	"io"
	"os"
)

type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return
	}
	_, sw.err = fmt.Fprintln(sw.w, s)
}
func wisdom(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	sw := safeWriter{w: f}
	sw.writeln("The road to wisdom? - Well, it's plain")
	sw.writeln("and simple to express:")
	sw.writeln("\tErr")
	sw.writeln("\tand err")
	sw.writeln("\tand err again,")
	sw.writeln("\tbut less")
	sw.writeln("\tand less")
	sw.writeln("\tand less.")
	sw.writeln("\t- Piet Hein")

	return sw.err
}

func main() {
	err := wisdom("wisdom.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
