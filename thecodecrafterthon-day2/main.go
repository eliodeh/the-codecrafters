package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("BASE CONVERTER")

	var num string
	var base string

	for {

		fmt.Println("enter any value")
		fmt.Scan(&num)
		if num == "quit" {
			fmt.Println("goodbye codecrafter")
			break
		}
		fmt.Println("enter any base converter: hex, bin, dec")
		fmt.Scan(&base)
		base = strings.ToLower(base)

		if base == "quit" {
			fmt.Println("goodbye codecrafter")
			break
		}
		if base == "hex" {
			value, err := strconv.ParseInt(num, 16, 64)
			if err != nil {
				fmt.Println("not a valid hex index")
			} else {
				fmt.Println(value)
			}
		}
		if base == "bin" {
			value, err := strconv.ParseInt(num, 2, 64)
			if err != nil {
				fmt.Println("not a valid bin index")
			} else {
				fmt.Println(value)
			}
		}

		if base == "dec" {
			num2, err := strconv.ParseInt(num, 10, 64)
			if err == nil {
				fmt.Println("not a valid dec value")
			} else {
				value1 := strconv.FormatInt(num2, 16)
				value2 := strconv.FormatInt(num2, 2)
				fmt.Println(value2)
				fmt.Println(strings.ToUpper(value1))
			}

		}

	}
}
