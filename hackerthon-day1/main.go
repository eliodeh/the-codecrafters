package main

import (
	"fmt"
)

func main() {
	const (
		colourred    = "\033[0;31m"
		colouryellow = "\033[0;33m"
		colourblue   = "\033[0;34m"
		colourpurple = "\033[0;35m"
		colourreset  = "\033[0m"
		colourgreen  = "\033[0;32m"
	)
	fmt.Println("CLI Calculator")
	for {
		var firstnum int
		var secondnum int
		var Operate string
		var digit string

		fmt.Println(colourred + "enter any firstnum" + colourreset)
		fmt.Scan(&firstnum)
		if _, err := fmt.Scan(&firstnum); err != nil {
			fmt.Println("invalid number")
			continue
		}
		fmt.Println(colourgreen + "enter any secondnum" + colourreset)
		fmt.Scan(&secondnum)
		if _, err := fmt.Scan(&secondnum); err != nil {
			fmt.Println("invalid number")
			continue
		}
		fmt.Println("enter any Operate add, sub, mul, div, help, quit")
		fmt.Scan(&Operate)
		if _, err := fmt.Scan(&Operate); err != nil {
			fmt.Println("invalid number")
			continue
		}
		fmt.Println("enter any digit")
		fmt.Scan(&digit)
		if _, err := fmt.Scan(&digit); err != nil {
			fmt.Println("invalid number")
		}

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
		fmt.Scan(Operate)
		fmt.Println(add(firstnum, secondnum))
		fmt.Println(sub(firstnum, secondnum))
		fmt.Println(mul(firstnum, secondnum))
		fmt.Println(div(firstnum, secondnum))
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
	if snum == 0 {
		fmt.Print("can't divide by 0")
	} else {
		return fnum / snum
	}
	return 3
}
