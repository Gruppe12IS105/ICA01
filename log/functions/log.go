package functions

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func LogaritmeFunc() {
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	arg2Float, _ := strconv.ParseFloat(arg2, 64)
	arg1Int, _ := strconv.Atoi(arg1)

	if arg1Int == 2 {
		fmt.Println(math.Log2(arg2Float))
	} else if arg1Int == 10 {
		fmt.Println(math.Log10(arg2Float))
	} else {
		fmt.Println("Første argumentet må være 2 eller 10")
	}
}
