package main

import "fmt"

type Item struct {
	name string
	price float64
}

func addCustomer(m map[string][]Item, name string){

}

func removeCustomer(m map[string][]Item, name string){
}

func addItem(m map[string][]Item, name string, item Item){

}

func printCustomerBill(m map[string][]Item, name string){

}

func printTotal(m map[string][]Item){

}

func main(){
	r := map[string][]Item{}
	addCustomer(r, "bob")
	addCustomer(r, "joe")
	addCustomer(r, "sally")
	addItem(r,"bob",Item{"pizza",12.3})
	addItem(r,"bob",Item{"coke",4.2})
	printCustomerBill(r,"bob")
	addItem(r,"sally",Item{"hotdog",5.9})
	addItem(r,"sally",Item{"salad",5.3})
	printCustomerBill(r,"sally")
	printTotal(r)
	removeCustomer(r,"sally")
	printCustomerBill(r,"sally")
	printTotal(r)
	
	
}