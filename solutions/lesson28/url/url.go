package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	u, err := url.Parse("https://a b.com/")

	if err != nil {
		fmt.Println(err)
		fmt.Printf("%#v\n", err)

		if e, ok := err.(*url.Error); ok {
			fmt.Println("Op:", e.Op)
			fmt.Println("URL:", e.URL)
			fmt.Println("Err:", e.Err)
		}
		os.Exit(1)
	}

	fmt.Println(u)
}
