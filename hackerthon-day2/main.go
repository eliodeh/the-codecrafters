package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("CLI calculator")

	var num string
	var base string

	for {

		fmt.Println("enter any num")
		fmt.Scan(&num)
		if num == "quit" {
			fmt.Println("goodbye codecrafter")
			break
		}
		fmt.Println("Enter the base")
		fmt.Scan(&base)
		if base == "quit" {
			fmt.Println("goodbye codecrafter")
			break
		}
		if base == "hex" {
			value, err := strconv.ParseInt(num, 16, 64)
			if err == nil {
				fmt.Println(value)
			}
			if err != nil {
				fmt.Println("not a valid hex index")

			}
		}
		if base == "bin" {
			value, err := strconv.ParseInt(num, 2, 64)
			if err == nil {
				fmt.Println(value)
			}
			if err != nil {
				fmt.Println("not a valid binary value")

			}
		}
		if base == "dec" {
			num2, err := strconv.ParseInt(num, 10, 64)
			if err == nil {
				value1 := strconv.FormatInt(num2, 16)
				value2 := strconv.FormatInt(num2, 2)
				fmt.Println(value1, value2)
			}
			if err != nil {
				fmt.Println("not a valid decimal value")

			}

		}
	}

}
