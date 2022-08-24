package main

import "fmt"

func main() {
	score := 94
	highScore := 99
	//Println() will always print in new line at the end of statement automatically
	fmt.Println("My score is " , score)

	//use format specifiers
	//use printf with format specifiers. cannot use PrintIn()

	//%v is one of the many format specifiers

	//Printf() will not print in new line at the end of statement automatically
	fmt.Printf("My score is %v\n ", score)
	fmt.Printf("My score is %v.\n ", score)
	fmt.Printf("My score is '%v'\n ", score)
	fmt.Printf("My score : %v '\n", score)

	//use multiple format specifiers in the single statement
	fmt.Printf("My score is %v. My highest score till date is %v\n",score,score)
	fmt.Printf("My score is %v. My highest score till date is %v\n",score,highScore)

	//%T => datatype
	fmt.Printf("value stores in score variable is %v. Datatype of score variable is %T\n",score,score)
	//working with space in format specifiers
	fmt.Printf("score is %10v.\n",score) //right align
	fmt.Printf("score is %-10v.\n",score) //left align
}