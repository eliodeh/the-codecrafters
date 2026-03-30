package main

import (
	"fmt"
	"strconv"
)

var firstnum string
var secondnum string
var Operate string

var val1 int
var val2 int
var opp string

// func main() {
// fmt.Println("CLI calculator")
func main() {
start:
	fmt.Println("enter any firstnum")
	fmt.Scan(&firstnum)
	v, err := strconv.Atoi(firstnum)
	if err == nil {
		val1 = v
	start1:
		fmt.Println("enter any secondnum")
		fmt.Scan(&secondnum)
		s, err := strconv.Atoi(firstnum)
		if err == nil {
			val2 = s
			fmt.Println("Operate add, sub, mul, div, pow, mod")
			fmt.Scan(&Operate)
			if Operate == "add" || Operate == "sub" || Operate == "mul" || Operate == "div" || Operate == "pow" || Operate == "mod" {
				Operate = opp
			} else {
				fmt.Println("Eerror")
			}
		} else {
			fmt.Println("Invalid Input")
			goto start1
		}
	} else {
		fmt.Println("Invalid Input")
		goto start
	}

	// fmt.Println("enter any secondnum")
	// fmt.Scan(&secondnum)
	// s, err := strconv.Atoi(firstnum)
	// if err == nil {
	// 	val2 = s
	// } else {
	// 	fmt.Println("Invalid Input")
	// }

	fmt.Println("Operate add, sub, mul, div, pow, mod")
	fmt.Scan(&Operate)

	if Operate == "add" {
		add()
	}
	// if Operate == "sub" {
	// 	sub()
	// }
	// if Operate == "mul" {
	// 	mul()
	// }
	// if Operate == "div" {
	// 	div()
	// }
	// if Operate == "mod" {
	// 	mod()
	// }
	// if Operate == "pow" {
	// 	pow()
	// }
	// goto start

}

func add() {
	fmt.Println(val1 + val2)
}

// func sub() {
// 	fmt.Println(firstnum - secondnum)
// }
// func mul() {
// 	fmt.Println(firstnum * secondnum)
// }
// func div() {
// 	fmt.Println(firstnum / secondnum)
// }
// func pow() {
// 	fmt.Println(math.Pow(float64(firstnum), float64(secondnum)))
// }
// func mod() {
// 	fmt.Println(firstnum/secondnum, "the remainder ==", firstnum%secondnum)
// }
