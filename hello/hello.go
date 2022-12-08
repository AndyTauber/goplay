package main

import (
	"fmt"
	"math"

	"github.com/AndyTauber/goplay/string_util"
)

func main() {
	var x string = "Hello, World"
	var year int = 2022
	fmt.Println(x, year)
	var next_year int32 = 2023
	fmt.Printf("type of x : %T\n", x)
	fmt.Printf("type of next_year : %T\n", next_year)
	fmt.Prinln(math.Sqry(5))
	fmt.Pringln(string_util.Reverse("reverse me"))
}
