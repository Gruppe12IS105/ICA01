package main

import "fmt"

func main() {
	fmt.Println("Hei\nVi er gruppe 12 og består av:\n")
	fmt.Println("Marius")
	fmt.Println("Nicolay was here")
	fmt.Println("Indra")
	fmt.Println("Hallvard")
	fmt.Println("Tommy")
	fmt.Println("Nicole")
	var navn = []byte("\x45\x6d\x69\x6c")
	fmt.Printf("%s\n", navn)

}
