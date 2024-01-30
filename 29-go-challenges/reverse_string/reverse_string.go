package main

import "fmt"

func main(){
	fmt.Println(reverseString("hello")) // olleh
	fmt.Println(reverseString("olleh")) // hello
}

func reverseString(str string) string{
	reverse := ""
	for i := 0; i < len(str); i++ {
		reverse = string(str[i]) + reverse
	}
	
	return reverse
}