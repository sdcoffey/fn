# Fn

Fn is a collection of functional golang utilities meant to assist with common
slice and map operations

## Installation

```bash
$ go get github.com/sdcoffey/fn`
```

## Reference

1. [Collections](#collections)
   1. [Select/Reject](#select-and-reject)
   1. [Map](#map)
   1. [Reduce](#reduce)
   1. [Zip](#zip)
   1. [First](#first)
   1. [Any](#any)
1. [Numbers](#numbers)
   1. [Sum](#sum)
   1. [Min/Max](#min-and-max)
1. [Utility](#utilities)
   1. [Must](#must)
   1. [Zero](#zero) 


## Collections

### Select and Reject

`Select` and `Reject` reduce a slice to only values that either satisfy or do not
satisfy a predicate.

```go
allInts := []int{1,2,3,4,5,6}

func isEven(number int, index int) bool {
    return item % 2 == 0
}

evens := fn.Select(allInts, isEven) // [2, 4, 6]
odds := fn.Reject(allInts, isEven)  // [1, 3, 5]
```

### Map

`Map` transforms values in a slice
```go
ints := []int{1,2,3}
strings := fn.Map(ints, func(number int, index int) string {
    return strconv.Iota(number)
})

fmt.Println(strings) // ["1", "2", "3"]
```

### Reduce

`Reduce` allows you to condense a slice into one end value

```go
combined := fn.Reduce([]string{"a", "b", "c"}, func(result, item string, index int) string {
	return result + item
}, "")

fmt.Println(combined) // "abc"
```

### Zip

`Zip` returns a set of key and value slices translated into a amp

```go
keys := []string{"one", "two", "three"}
values := []int{1, 2, 3}

fmt.Println(Zip(keys, values)) // map[string]int{ "one": 1, "two": 2, "three": 3 }
```

### First
`First` returns the first value and index of a slice that satisfies a predicate, or the zero value and -1

```go
values := []int{-2,-1,0,1,2}
first, index := fn.First(values, func(item, index int) bool {
	return item > 0
})

fmt.Printf("%d at index %d", first, index) // 1 at index 3
```

### Any

`Any` returns true if any value in a slice satisifies a predicate

```go
ints := []int{2, 3}

fmt.Println(Any(ints, func(item, index int) bool {
   return item%2 == 1
})) // true

fmt.Println(Any(ints, func(item, index int) bool {
    return item == 10
})) // false
```

### AnyNonZero

`AnyNonZero` is a helper that returns true if any value in a slice is not the zero value

```go
fmt.Println(fn.AnyNonZero([]{0, 0, 0}))           // false
fmt.Println(fn.AnyNonZero([]{"", "", ""}))        // false
fmt.Println(fn.AnyNonZero([]{"", "abcd", ""}))    // true

type Example struct {
	Name string
}

fmt.Println(fn.AnyNonZero([]Example{{Name: "abcd"}, {}})) // true
```

## Numbers

### Min and Max

`Max` returns the max of some orderable slice
`Min` returns the min of some orderable slice

```go
strings := []string{"one", "two", "three"}
ints := []int{1, 2, 3}

fmt.Println(fn.Max(strings...))    // "two"
fmt.Println(fn.Max(ints...))       // 3

fmt.Println(fn.Min(strings...))    // "one
fmt.Println(fn.Min(ints...))       // 1
```

### Sum

`Sum` returns the Sum of some number value

```go
fmt.Println(fn.Sum([]int{1,2,3}))               // 6
fmt.Println(Sum([]float64{math.Pi, math.Pi}))   // 6.283185307179586
```

## Utilities

### Must

Get the value from an error-returning method, or panic

```go
func CanFail(shouldFail bool) (string, error) {
    if shouldFail {
        return "", fmt.Errorf("failed")
    }
    return "success", nil
}

val1 := fn.Must(CanFail(true)) // "success"
val2 := fn.Must(CanFail(false)) // panics
```

### Zero

`Zero` returns true if the value provided is the zero value for the type

```go
fmt.Println(fn.Zero(0)) // true
fmt.Println(fn.Zero(1)) // false

type Example struct {
    Name string
}

fmt.Println(fn.Zero(Example{}))             // true
fmt.Println(fn.Zero(Example{Name:"abcd"}))  // false

```
