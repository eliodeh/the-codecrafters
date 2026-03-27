package main

import (
	"fmt"
)

func main() {
	fmt.Println("CLI Calculator")
	for {
		var firstnum float64
		var secondnum float64
		var Operate string

		start:
		fmt.Println("enter any firstnum")
		_, err1 := fmt.Scan(&firstnum)
		fmt.Println("enter any secondnum")
		_, err2 := fmt.Scan(&secondnum)

		if err1 != nil || err2 != nil {
			fmt.Println("Input an valid  number")
			goto start
		}
		fmt.Println("enter any Operate add, sub, mul, div, help, quit")
		fmt.Scan(&Operate)

		if Operate == "add" {
			fmt.Println(add(firstnum, secondnum))
		}
		if Operate == "sub" {
			fmt.Println(sub(firstnum, secondnum))
		}
		if Operate == "mul" {
			fmt.Println(mul(firstnum, secondnum))
		}
		if Operate == "div" {
			fmt.Println(div(firstnum, secondnum))
		}
		if Operate == "help" {
			fmt.Println("add, sub, mul, div, help, quit")
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
		fmt.Print("can't divide by ")
	} else {
		return fnum / snum
	}
	return 0
}
