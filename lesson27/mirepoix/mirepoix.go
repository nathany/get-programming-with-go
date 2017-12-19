package main

import "fmt"

func main() {
	soup := mirepoix(nil)
	fmt.Println(soup)
}

func mirepoix(ingredients []string) []string {
	return append(ingredients, "onion", "carrot", "celery")
}
