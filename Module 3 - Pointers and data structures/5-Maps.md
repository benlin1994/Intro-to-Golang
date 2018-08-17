# Maps

## Maps

Maps are data structures that hold key value pairs.

## Declaring a map

A map can be declared using the following syntax:

```go
var m map[string] int
```

The key type is specified in the brackets while the value type is specified after the brackets.

An unitialized map has a value of `nil`.

## Initialing a map with a map literal

A map can be initialized using a map literal:

```go
    m := map[string] int{} //initializes an empty map
```

Key-value pairs can be initialized as well:

```go
    m := map[string] int{"a": 1, "b": 2, "c": 3} 
```

The initialized key-value pairs must be of the correct type.

## Accessing and modifying map values

Map values can be accessed and modified through their keys:

```go
    m := map[string] int{"a": 1, "b": 2, "c": 3} 
    fmt.println(m["a"]) //prints 1
    map["a"] = 111 
    fmt.println(m["a"]) //prints 111
```

## Deleting map key-value pairs

Map key-value pairs can be deleted from the map:

```go
    m := map[string] int{"a": 1, "b": 2, "c": 3} 
    delete(m, "a") //deletes key "a" from map
    fmt.println(m) //prints map[b:2 c:3]
```

## Testing if a key is present

If a key is not present in a map, accessing that key will result in the zero value of the map key type:

```go
    m := map[string] int{"a": 1, "b": 2, "c": 3} 
    val := m["d"] // "d" does not exist as a key
    fmt.println(val) //prints "", the zero value of a string
```

In the above scenario, we can not tell whether the key was not present or if the value just happened to be the same as the value type zero value.

However, `map[key]` conveniently has a second optional return value that returns true when the key is present and false when the key is not present:

```go
    m := map[string] int{"a": 1, "b": 2, "c": 3} 
    val, ok := m["d"]
    fmt.println(val, ok) //prints 0 false
```
