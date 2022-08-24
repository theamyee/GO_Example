package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//fmt.Println(os.Args)
	fmt.Println(os.Args[1:])
	//fmt.Println(reflect.TypeOf(os.Args[1]))
	sum := 0
	for _, v := range os.Args[1:] {
		//sum = sum + v
		//value, error := strconv.Atoi(v)
		value, _ := strconv.Atoi(v)
		// if error != nil {
		// 	fmt.Println(error)
		// 	os.Exit(1)
		// }
		sum = sum + value

	}
	fmt.Println("Sum of all numbers passed as arguments: ", sum)
}
