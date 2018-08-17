# Slices

## Slices

Slices are references that point to a specific section of an array. While arrays have a fixed length, slices have a dynamically sized length. Futhermore, modifying a slice results in modifying the underlying array.

## Declaring a slice


Declaring a slice is similar to declaring an array, except with no length specified:

```go
    var s []int
```

Unitialized slices have a value of `nil`.

## Initializing a slice with an array

Slices can be initialized by specifying a subsection of an array:

```go
    arr := [10]int{1,2,3,4,5,6,7,8,9,10}
    var s []int = arr[2:5] //slice from 2nd to 5th index (not including 5th)
    s2 := arr[0:2] //short declaration works too
    fmt.Println(s) //prints [3 4 5]
    fmt.Println(s2) //prints [1 2]
```

The array subsection that the slice references is specified by a start and end index. The slice will contain values starting from the start index and ending with the value right before the end index. You can not use negative slice indexes.

## Initializing a slice with a slice literal

Slices can be initialized with a slice literal:

```go
	s := []int{1,2,3,4,5}
```

Behind the scenes, Go creates an array and then has the slice reference it.


## Slice length and capacity

Slices have a length and capacity type that can be accessed through the methods len(slice) and cap(slice):

```go
    s := []int{1,2,3,4,5}
    fmt.Println(len(s, cap(s)))//prints 5 5
```

The slice length represents how many values the slice contains. The slice cap represents how values the underlying array contains.

The slice length can be extended through reslicing but can never surpass the cap limit.

## Reslicing slices

Slices can also slice from other slices:

```go
	s := []int{1,2,3,4,5}
    s = s[0:3] // contains [1 2 3], len = 3, cap = 5
    s = s[0:5] // contains [1 2 3 4 5], len 5, cap = 5
```

The slice end index can be decreased and increased to contain a different subsection of the original array.

However, increasing the slice start index will change the start pointer and decrease the slice cap. Decreasing the start index again won't be able to undo that:

```go
	s := []int{1,2,3,4,5}
    s = s[2:5] // contains [3 4 5] , len = 3, cap = 3
    s = s[0:5] // can't do this! out of bounds 
```

## Slice defaults

If the start or end index is omitted from a slice initialization, the index will default to the 0th or last index of the array.

```go
    s := []int{1,2,3,4,5}
    s = s[:5] // slice from 0th to 5th index (not including 5th)
    s = [2:] // slice from 2nd to last index
    s = [:] //slice from 0th to last index
```

## Appending to a slice

The `append()` method can be used to append an element to a slice:

```go
    s := []int{1,2,3} // contains [1 2 3]
    s.append(4) // contains [1 2 3 4]
```

Multiple values can be appended at the same time:

```go
    s := []int{1,2,3} // contains [1 2 3]
    s.append(4,5,6) // contains [1 2 3 4 5 6]
```

