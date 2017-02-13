package main

import (
	"flag"
	"os"
)

func main() {

	argsA := os.Args[1:]
	argsB := os.Args[1:]

	var argsAsomInt = flag.Int("argsA", argsA, "")

	var argsBsomInt = flag.Int("argsB", argsB, "Buhu")
	flag.Parse()
	return argsAsomInt
}

/*
	argsWithoutProg1 := os.Args[1:]
	arg := os.Args[2]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
*/
