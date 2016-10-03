package main

import "fmt"

func rev(str string) {
	l:=len(str)
	if(l==0){

	}else {
		fmt.Print(string(str[l-1]))
		str1:=str[:l-1]
		rev(str1)
	}
}
func main() {
	fmt.Println("Enter  a string")
	var str string
	fmt.Scanf("%s",&str)
	fmt.Println(str)
	rev(str)
}
