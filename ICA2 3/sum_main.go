package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	argsA := os.Args[1]
	argsB := os.Args[2]
	a, _ := strconv.ParseFloat(argsA, 64)
	b, _ := strconv.ParseFloat(argsB, 64)

	fmt.Println(a + b)
}

/*
	argsWithoutProg1 := os.Args[1:]
	arg := os.Args[2]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
*/
