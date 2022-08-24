package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	//basic function call with no return values
	welcome()

	//functions can return values
	name := "jio"
	check := welcomeByName(name)
	if check {
		fmt.Println("Welcome,", name)
	} else {
		fmt.Println(name, ", your name is not in the guest list!!!!")
	}

	

	//... means that give this array the size as per count of value
	scores := []int{10, 20, 30, 40}
	length, total := getSizeAndCount(scores)
	fmt.Println("Length: ", length)
	fmt.Println("Total: ", total)

}

// func getSizeAndCount(s []int) (int, int) {
// 	size := len(s)
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	return size, sum
// }


func getSizeAndCount(s []int) (size int, sum int) {
	size = len(s)
	sum = 0
	for _, v := range s {
		sum += v
	}
	return
}

func second(){
	name := os.Args[1]
	check := welcomeByName(name)
	if check {
	fmt.Println("Welcome, ",name)
	} else{
	fmt.Println(name, ", your name is not in the list! ")
	}
}

func welcome() {
	fmt.Println("Welcome to Uranus!")
}

func welcomeByName(nm string) bool {
	result1 := strings.Contains(nm, "a")
	result2 := strings.Contains(nm, "e")
	result3 := strings.Contains(nm, "i")
	result4 := strings.Contains(nm, "o")
	result5 := strings.Contains(nm, "u")

	//return true or return false
	//if any of the result1, result2, result3, result4, result5 is
	//true then return true or return false
	if result1 || result2 || result3 || result4 || result5 {
		return true
	} else {
		return false
	}
}