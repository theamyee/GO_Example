// array => collection of the values of same data type

package main

import (
	"fmt"
	"reflect"
)

func main(){

	//declare array
	var scores = [8]int{4,6,8,10,12,14,16,18}
	fmt.Println(scores)

	var friends = [3]string{"mary","alex","john"}
	fmt.Println(friends)

	var specificPosition = [9]int{0:7,1:22,2:6,3:4,4:66}
	fmt.Println(specificPosition)

	fmt.Printf("The length of the array specificPosition is %v\n",len(specificPosition))
	fmt.Printf("The data type of the array specificPosition is %T\n",specificPosition)
	fmt.Println("Datatype specificPosition is ", reflect.TypeOf(specificPosition).Kind())


}