package main

import (
	"fmt"
	"os"
)

func wisdom(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(f, "The road to wisdom? - Well, it's plain")
	if err != nil {
		f.Close()
		return err
	}

	_, err = fmt.Fprintln(f, "and simple to express:")
	f.Close()
	return err
}

func main() {
	err := wisdom("wisdom.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
