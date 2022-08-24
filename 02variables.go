package main

import (
	"fmt"
	"reflect"
)

func main() {
	//declare the variable
	//var variableName datatype
	//declare variable score
	var score int
	//assign value 8 to variable score
	//initialise variable score with value 8
	score = 8
	score = score + 1
	//if you declare the variables,then use them
	fmt.Println("Go is strictly typed language!")
	fmt.Println(score)

	//check the datatype of variable, use reflect package
	fmt.Println(reflect.TypeOf(score).Kind())

	//string variable
	var message string
	message = "Go lang makes beginner developers good developers!"
	fmt.Println(message)
	fmt.Println(reflect.TypeOf(message))

	//float values
	var lifeline float32
	lifeline = 44.44
	fmt.Println(lifeline)
	fmt.Println(reflect.TypeOf(lifeline))

	//boolean values
	var check bool
	check = true
	fmt.Println(check)
	fmt.Println(reflect.TypeOf(check))

	//auto type inference
	//use := for auto type inference
	//life is not yet declared
	//the use of := for auto type inference
	fmt.Println("using auto type inference!")
	life := 90
	fmt.Println(life)
	fmt.Println(reflect.TypeOf(life))

	myMessage:= "hello"
	fmt.Println(myMessage)
	fmt.Println(reflect.TypeOf(myMessage))

	age:= 71.71
	fmt.Println(age)
	fmt.Println(reflect.TypeOf(age))

}