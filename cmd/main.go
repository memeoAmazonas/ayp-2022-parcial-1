package main

import (
	"fmt"
	"github.com/memeoAmazonas/ayp-2022-parcial-1/internal"
)

func main() {

	var values = []string{"1", "2", "3", "4", "5"}
	var values3 = []string{"1", "a", "3", "n", "a", "0"}
	internal.Response(values)
	fmt.Println(internal.Exercise_2(-2, 6, 2))
	fmt.Println(internal.Exercise_3(values3))
}
