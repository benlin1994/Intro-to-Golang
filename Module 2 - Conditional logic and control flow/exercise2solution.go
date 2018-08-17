package main

import "fmt"

func printCheckboard(){
	var one = "X"
	var two = "O"
	for i:= 0; i < 8; i++ {
		for j:=0; j < 8; j++{
			if j % 2 == 0{
				fmt.Print(one)
			}else{
				fmt.Print(two)
			}
		}
		fmt.Println()
		if one == "X"{
			two = "X"
			one = "O"
		} else if one == "O"{
			one = "X"
			two = "O"
		}
			
	}
}
func main(){
	printCheckboard()
}