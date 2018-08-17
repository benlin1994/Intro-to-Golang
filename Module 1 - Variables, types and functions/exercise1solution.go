package main

import "fmt"


//radiusToCirc
func radiusToCirc(rad float64) float64{
	const Pi = 3.14
	return Pi * rad * 2
}

//isShorter
func isShorter(str string, num int) bool{
	return len(str) <= num
}
//swap
func swap(x,y int)(a,b int){
	a = y
	b = x
	return
}

	
func main(){

	fmt.Println(radiusToCirc(4.2))
	fmt.Println(radiusToCirc(0))
	fmt.Println(isShorter("longer",2))	
	fmt.Println(isShorter("short",10))	
	fmt.Println(isShorter("short",5))	
	a,b := swap(4,-5)
	fmt.Println(a)
	fmt.Println(b)
}