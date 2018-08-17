# Arrays

## Arrays
Arrays hold a list of values. Arrays are declared using the following syntax:

```go
var a[5] int
```

Unitialized arrays have their values initialized to their type's zero values. Arrays have their length defined upon initialization. Arrays can not be resized and can only hold up to their length.




## Accessing and modifying array values

Array values can be accessed and modified through their index:

```go
var a[5] int
a[0] = 42
fmt.Println(a[0]) //prints 42
```

## Arrays as parameters

In Go, arrays are passed by value in parameters and not by reference. Modifying an array passed as a parameter won't change the original array. You must also specify the array length in the parameter.


```go
array := [5]int{1,2,3,4,5}

func modifyArray(arr [5]int){
  arr[0] = 100
}
modifyArray(array)
fmt.Println(array[0]) // still 1
```

## Initializing arrays with array literals

Array values can be initialized using an array literal. To use an array literal, use the following syntax:

```go
array := [5]int{1,2,3,4,5}
fmt.Println(array)// prints [1 2 3 4 5]
```

## Using range to iterate through an array

The `range` statement can be used to in a `for` loop to iterate through an array:

```go
array := [5]string{"a","b","c","d","e"}

for i, val := range array{
    fmt.Printf("index: %d | value: %s\n", i, val)
}
//prints index: i, value: val for each element in the array
```

If you only want the index and not the value, drop the `val` variable:

```go
array := [5]string{"a","b","c","d","e"}

for i := range array{
    fmt.Printf(array[i])
}
```

If you only want the value, but not the index, replace the `i` with a _:

```go
array := [5]string{"a","b","c","d","e"}

for _, val := range array{
    fmt.Printf(val)
}
```