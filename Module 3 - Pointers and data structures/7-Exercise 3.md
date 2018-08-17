# Module 3 Exercise


## Instructions

Download the Module 3 exercise file named `exercise3.go`. In the main function, there is a map `r` that represents a restaurant. Map `r` has customer names as strings for keys and slices of Item types for values. Item types are structs that have a name and a price field. 

The main function executes a series of functions that modify the map `r`. Your exercise is to implement all of the following functions so that they pass all of the test cases.

* func addCustomer(m map[string][]Item, name string) - adds the name string as a key to the map. The value should be an empty Item slice.
* func removeCustomer(m map[string][]Item, name string) - removes the name string as a key to the map.
* func addItem(m map[string][]Item, name string, item Item) - adds an Item to the Item slice of the key specified by the name string.
* func printCustomerBill(m map[string][]Item, name string) - adds up the bill for a single customer and prints it
* func printTotal(m map[string][]Item) - adds up the bill for all customers in the restaurant and prints it


## Test your results

Test your results by running `exercise3.go` with the `go run exercise3.go` command. If your program outputs the following, it should be working:

```
bob : $16.500000
sally : $11.200000
Total: $27.700000
sally : $0.000000
Total: $16.500000
```