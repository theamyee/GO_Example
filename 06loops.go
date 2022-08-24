package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Printf("20 x %v = %v\n",i,i*20)
	}

	fmt.Println("-------------------------------")
	var scores = [8]int{4,6,8,10,12,14,16,18}
	//iterate through the stored array
	for index, value := range scores{
		fmt.Printf("Index: %v, value: %v \n", index, value)
	}
	fmt.Println("-------------------------------")
	//iterate through the scores array
	for _, value := range scores {
		fmt.Printf("Value: %v \n", value)
	}

}