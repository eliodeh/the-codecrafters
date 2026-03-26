package main

import "fmt"

func main() {
	fmt.Println("CLI Calculator")
	for {
		var firstnum int
		var secondnum int
		var Operate string

		fmt.Println("enter any firstnum")
		fmt.Scan(&firstnum)
		fmt.Println("enter any secondnum")
		fmt.Scan(&secondnum)
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
			fmt.Println("Help ")
		}
		if Operate == "quit" {
			fmt.Println("Goodbye codecrafterthon")
			break
		}
	}
}

// the running functions of the maths //
func add(fnum, snum int) int {
	return fnum + snum
}
func sub(fnum, snum int) int {
	return fnum - snum
}
func mul(fnum, snum int) int {
	return fnum * snum
}
func div(fnum, snum int) int {
	return fnum / snum
}
