# Structs

## Structs

Structs are a collection of fields. They are the closest thing Go has to classes or objects.

Use the following syntax to define a struct type:

```go
type Point struct {
	X int
	Y int
}
```


## Declaring a struct

Use the following syntax to declare a struct:

```go
    point := Point{} // point is of type Point
```

Uninitialized struct fields are initialized to their type's zero value.

## Accessing and modifying the fields of a struct

Struct fields can be accessed and modified with the `.` notation:

```go
    point := Point{}  
    fmt.Println(point.X) //prints 0
    fmt.Println(point.Y) //prints 0
    point.X = 2
    point.Y = 4
    fmt.Println(point.X) //prints 2
    fmt.Println(point.Y) //prints 4
```

## Initializing a struct with struct literals

Struct literals are used to initialize a struct's fields. They are used with the following syntax:

```go
    point := Point{2, 4}  // v1 has type Point, 
    fmt.Println(point.X) //prints 2
    fmt.Println(point.Y) //prints 4
```

To fully initialize using struct literal, a value is required for each field in the struct.The values have to be in the same order as the fields and be of the correct field type.

## Initializing specific fields of a struct

If you only want to initialize some fields and not all, you can specify which ones you want by using their field name:

```go
    point := Point(Y:2) // initializes point.Y to 2
    fmt.Println(point.X) //prints 0
    fmt.Println(point.Y) //prints 2
```

## Struct pointers

Struct types can have pointers as well:

```go
    p := &Point(2,4) //p is of type *Point
```

Struct fields can still be accessed from pointers using the `.` notation:


```go
    p := &Point(2,4) //p is of type *Point
    fmt.Println(p.X) //prints 0
    fmt.Println(p.Y) //prints 2
    p.X = 3
    p.Y = 5
    fmt.Println(p.X) //prints 3
    fmt.Println(p.Y) //prints 5
```

Behind the scenes, Go is actually dereferencing the pointer for us:

```go
    p := &Point(2,4) 
    p.X = 3 //behind the scenes: *(p).X = 3
```