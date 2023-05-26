package main

import (
	"fmt"
)

var Val1 int = 100 //global level ---(pascal case)

var myValue int = 500 //package level –(camel case)
func main2() {
	var val1 = 55 //local variable

	fmt.Println(val1)
	fmt.Println(myValue)
}
