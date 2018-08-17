# Pointers


## Pointers
Pointers hold the memory address of a value.

To create a pointer type, use the `*` symbol:


```go
var p *int //p is of type * int
```

Uninitialized pointers have a zero value of `nil`

```go
var p *int 
fmt.println(p)
//prints <nil>
```

## & operator
To get the memory address of a value, use the `&` operator and assign the result to a pointer type:

```go
i:=123 //i of type int
var p * int = &i //p is of type *int
q := &i //short declaration is okay

```

## * operator
To get the value at a pointer to a memory address, use the `*` operator:

```go
i ;= 123
p := &i
fmt.Println(p) // prints some memory address
fmt.Println(*p) // prints 123
```


You can also edit the value of a pointer using the `*` operator:

```go
i ;= 123
p := &i
*p = 456
fmt.Println(*p) // prints 456
```

Accessing a pointers value using the `*` operator is referred to as dereferencing. 

## Go has no pointer arithmetic

Go has no pointer arithmetic. You can not add to pointers to get to different memory addresses:

```go
i ;= 123
p := &i
*(p+8) = 456 //can't do this!
```