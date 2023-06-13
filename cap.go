package main

import (
	"fmt"
	"strings"
)

func main() {
	var txt string
	var res1 string
	fmt.Println("enter the text: ")
	fmt.Scanln(&txt)
	fmt.Println("before capitalize:", txt)
	res1 = strings.ToUpper(txt)
	fmt.Println("result:", res1)

}
