package main

import "fmt"

type Item struct {
	name string
	price float64
}


func addCustomer(m map[string][]Item, name string){
	if _, ok := m[name]; !ok{
		m[name] = make([]Item,10)
	}	
}

func removeCustomer(m map[string][]Item, name string){
	delete(m,name)
}

func addItem(m map[string][]Item, name string, item Item){
	if val, ok := m[name]; ok{
		m[name] = append(val,item)
	}	
}

func printCustomerBill(m map[string][]Item, name string){
	sum := 0.0
	
	if val, ok := m[name]; ok {
		for _, item := range val {
			sum += item.price
		}
	}	
	fmt.Printf("%s : $%f\n",name,sum)
}

func printTotal(m map[string][]Item){
	sum := 0.0
	
	for _, val := range m {
		for _, item := range val {
			sum += item.price
		}
	}
	fmt.Printf("Total: $%f\n",sum)
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