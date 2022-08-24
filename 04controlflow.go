package main

import "fmt"

func main() {

	score := 88
	if score > 88 {
		fmt.Printf("%v is greater than 88.\n",score)
	}

	if score >= 90{
		fmt.Printf("%v is greater than 90.\n",score)
	}else{
		fmt.Printf("%v is greter than 90.\n",score)
	}
}