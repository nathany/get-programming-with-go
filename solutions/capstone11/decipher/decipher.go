package main

import "fmt"

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	message := ""
	keyIndex := 0

	for i := 0; i < len(cipherText); i++ {
		// A=0, B=1, ... Z=25
		c := cipherText[i] - 'A'
		k := keyword[keyIndex] - 'A'

		// cipher letter - key letter
		c = (c-k+26)%26 + 'A'
		message += string(c)

		// increment keyIndex
		keyIndex++
		keyIndex %= len(keyword)
	}

	fmt.Println(message)
}
