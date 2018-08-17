# Basics

## Types

Go has the following basic types:

* int, int8, int16, int32, int64 
* uint, uint8, uint16, uint32, uint64, uintptr
* float32, float64 
* string 
* bool
* byte - alias for uint8
* rune - alias for int32 and used for Unicode characters
* complex64, complex128 - used for complex numbers

`int`, `uint` and `uintptr` are either 32 bit or 64 bit depending on the operating system.

## Comparison
Go has the following symbols used for comparison:

* `==` - equal
* `!=`- not equal
* `<` - less than
* `<=` - less than or equal
* `>` - greater than
* `>=` - greater than or equal

## Logical operators

Go has the following symbols used for logical operators:

* `&&` - logical and
* `||` - logical or
* `!` - logical not

## Arithmetic

Go has the following symbols use for arithmetic: 

* `+` - addition 
* `-` - subtraction
* `*` - multiplication
* `/` - quotient
* `%` - remainder
* `&` - bitwise and
* `|` - bitwise or
* `^` - bitwise xor
* `&^` - bit clear 
* `<<` - left shift
* `>>` - right shift

## Pointers

Go has the following symbols used for pointers:

* `&` - address of a pointer
* `*` - deference a pointer

## Channels

Go has the following symbol used for channels:

* `<-` send/receive data from a channel

## Nil

Go has a `nil` value that represents a zero value for uninitialized pointers, maps, slices, interfaces, functions and channels 