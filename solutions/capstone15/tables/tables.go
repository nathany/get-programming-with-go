package main

import (
	"fmt"
)

// Celsius temperatures
type Celsius float64

// Fahrenheit converts ºC to ºF.
func (c Celsius) Fahrenheit() Fahrenheit {
	return Fahrenheit((c * 9.0 / 5.0) + 32.0)
}

// Fahrenheit temperatures
type Fahrenheit float64

// Celsius converts ºF to ºC.
func (f Fahrenheit) Celsius() Celsius {
	return Celsius((f - 32.0) * 5.0 / 9.0)
}

const (
	line         = "======================="
	rowFormat    = "| %8s | %8s |\n"
	numberFormat = "%.1f"
)

type getRowFn func(row int) (string, string)

// drawTable draws a two column table.
func drawTable(hdr1, hdr2 string, rows int, getRow getRowFn) {
	fmt.Println(line)
	fmt.Printf(rowFormat, hdr1, hdr2)
	fmt.Println(line)
	for row := 0; row < rows; row++ {
		cell1, cell2 := getRow(row)
		fmt.Printf(rowFormat, cell1, cell2)
	}
	fmt.Println(line)
}

func ctof(row int) (string, string) {
	c := Celsius(row*5 - 40)
	f := c.Fahrenheit()
	cell1 := fmt.Sprintf(numberFormat, c)
	cell2 := fmt.Sprintf(numberFormat, f)
	return cell1, cell2
}

func ftoc(row int) (string, string) {
	f := Fahrenheit(row*5 - 40)
	c := f.Celsius()
	cell1 := fmt.Sprintf(numberFormat, f)
	cell2 := fmt.Sprintf(numberFormat, c)
	return cell1, cell2
}

func main() {
	drawTable("ºC", "ºF", 29, ctof)
	fmt.Println()
	drawTable("ºF", "ºC", 29, ftoc)
}
