package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("CLI Calculator")
	for {
		var input1 string
		var input2 string
		var Operate string

		fmt.Println("enter any firstnum")
		fmt.Scan(&input1)
		num1, err1 := strconv.ParseFloat(input1, 64)
		if err1 != nil {
			fmt.Println("error:invalid input")
			continue
		}

		fmt.Println("enter any secondnum")
		fmt.Scan(&input2)
		num2, err2 := strconv.ParseFloat(input2, 64)
		if err2 != nil {
			fmt.Println("invalid input")
			continue
		}

		fmt.Println("enter any Operate add, sub, mul, div, help, quit")
		fmt.Scan(&Operate)

		if Operate == "add" {
			fmt.Println(add(num1, num2))
		}
		if Operate == "help" {
			fmt.Println("Help ")
			fmt.Println("add, sub, mul, div, help, quit")
		}
		if Operate == "div" {
			fmt.Println(div(num1, num2))
		}
		if Operate == "mul" {
			fmt.Println(mul(num1, num2))
		}
		if Operate == "sub" {
			fmt.Println(sub(num1, num2))
		}
		if Operate == "quit" {
			fmt.Println("Goodbye codecrafterthon")
			break
		}
	}
}

// the running functions of the maths //
func add(fnum, snum float64) float64 {
	return fnum + snum
}

func sub(fnum, snum float64) float64 {
	return fnum - snum
}

func mul(fnum, snum float64) float64 {
	return fnum * snum
}

func div(fnum, snum float64) float64 {
	if snum == 0 {
		fmt.Print("can't divide by 0")
	} else {
		return fnum / snum
	}
	return 0
}
