package main

import "./boring"

//"fmt"

func main() {
	msg := "Hallo" //os.Args[1]
	boring.Boring01(msg)

	//Metoden nedenfor er for å kjøre Boring10
	/*
	  c := make(chan string) //oppretter channel c i denne filen

	  go boring.Boring10("Index nr: ", c) //starter func Boring10 i den andre koden
	  for i := 0; i < 1; i++{             //denne for-løkken kjører parallellt med Boring10
	    fmt.Printf("Test er: %q\n", <-c)
	  }
	  fmt.Println("Neiass, dette gidder jeg ikke")
	*/
}
