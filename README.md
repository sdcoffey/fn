# fn

Package fn is a collection of functional golang utilities meant to assist with
common slice and map operations, as well as provide some better dexterity around
error handling.

## Installation

```
go get github.com/sdcoffey/fn
```
## In this package
* [Any](#Any)
* [AnyNonZero](#AnyNonZero)
* [First](#First)
* [GenSeq](#GenSeq)
* [Map](#Map)
* [Max](#Max)
* [Min](#Min)
* [Must](#Must)
* [Reduce](#Reduce)
* [Reject](#Reject)
* [Select](#Select)
* [Seq](#Seq)
* [Sum](#Sum)
* [Zero](#Zero)
* [Zip](#Zip)


### Any
Any returns true if any item in the slice satisfies the predicate

```go
func ExampleAny() {
	negativeInts := []int{-1, -2, -3}
	anyPositive := fn.Any(negativeInts, func(item int, index int) bool {
		return item > 0
	})

	fmt.Println(anyPositive)
	// Output: false
}
```

### AnyNonZero
AnyNonZero is a helper function that returns true if any item in the slice
is not the zero value for its type

```go
func ExampleAnyNonZero() {
	nonZeroStrings := []string{"one", "", "", "four"}
	fmt.Println(fn.AnyNonZero(nonZeroStrings))

	zeroStrings := []string{"", ""}
	fmt.Println(fn.AnyNonZero(zeroStrings))
	// Output:
	// true
	// false
}
```

### First
First will return the first value that satisfies the predicate, and the
index at which the value was found or -1 if it was not found

```go
func ExampleFirst() {
	sequence := []int{-1, 0, 1, 2}
	firstPositive, found := fn.First(sequence, func(item int, index int) bool {
		return item > 0
	})

	fmt.Println("Value found:", found)
	fmt.Println("First positive value:", firstPositive)
	// Output:
	// Value found: 2
	// First positive value: 1
}
```

### GenSeq
GenSeq returns a chan that yields integers from [start, end), incremented by
inc

```go
func ExampleGenSeq() {
	intChan := fn.GenSeq(10, 0, -2)

	for value := range intChan {
		fmt.Print(value, " ")
	}
	// Output: 10 8 6 4 2
}
```

### Map
Map transforms allows you to transform all values into a slice of other
values. The original slice is not modified.

```go
func ExampleMap() {
	ints := []int{1, 2, 3}
	doubled := fn.Map(ints, func(item int, index int) int {
		return item * 2
	})

	fmt.Println(doubled)
	// Output: [2 4 6]
}
```

### Max
Max returns the maximum value in some orderable set, or the zero value of
the set type if it has len 0

```go
func ExampleMax() {
	fmt.Println(fn.Max(1, 2, 100, -1))
	// Output: 100
}
```

### Min
Min returns the minimum value in some orderable set, or the zero value of
the set type if it has len 0

```go
func ExampleMin() {
	fmt.Println(fn.Min(-100, -300, 3, 100))
	// Output: -300
}
```

### Must
Must allows you to return one value from a function that would normally
return a value and an error. If the error is present, Must will panic

```go
func ExampleMust() {
	canFail := func(shouldFail bool) (string, error) {
		if shouldFail {
			return "", fmt.Errorf("error")
		}
		return "success", nil
	}

	value := fn.Must(canFail(false))

	fmt.Println(value)
	// Output: success
}
```

### Reduce
Reduce condenses a collection into an accumulated value which is the result
of running each element in the collection though `reducer`, where each
successive iteration is supplied the return value of the previous.

```go
func ExampleReduce() {
	combined := fn.Reduce([]string{"a", "b", "c"}, func(result, item string, index int) string {
		return result + item
	}, "")

	fmt.Println(combined)
	// Output: abc
}
```

### Reject
Reject reduces a slice to only values for which `rejector` returns false

```go
func ExampleReject() {
	allInts := []int{1, 2, 3, 4, 5, 6}

	isEven := func(number int, index int) bool {
		return number%2 == 0
	}

	oddNumbers := fn.Reject(allInts, isEven)

	fmt.Println(oddNumbers)
	// Output: [1 3 5]
}
```

### Select
Select reduces a slice to only values for which `selector` returns true

```go
func ExampleSelect() {
	allInts := []int{1, 2, 3, 4, 5, 6}

	isEven := func(number int, index int) bool {
		return number%2 == 0
	}

	evenNumbers := fn.Select(allInts, isEven)

	fmt.Println(evenNumbers)
	// Output: [2 4 6]
}
```

### Seq
Seq Returns a slice of integers from [start, end), incremented by inc

```go
func ExampleSeq() {
	seq := fn.Seq(0, 50, 5)

	fmt.Println(seq)
	// Output: [0 5 10 15 20 25 30 35 40 45]
}
```

### Sum
Sum returns the summed value of values in the slice

```go
func ExampleSum() {
	fmt.Println(fn.Sum([]int{1, 2, 3}))
	fmt.Println(fn.Sum([]float64{math.Pi, math.Pi}))
	// Output:
	// 6
	// 6.283185307179586
}
```

### Zero
Zero returns whether the comparable value passed in is the zero value for
its type

```go
func ExampleZero() {
	fmt.Println(fn.Zero(""))
	fmt.Println(fn.Zero(1))

	type Example struct {
		Name string
	}

	fmt.Println(fn.Zero(Example{}))
	fmt.Println(fn.Zero(Example{Name: "abcd"}))

	// Output:
	// true
	// false
	// true
	// false
}
```

### Zip
Zip takes two slices and aggregates them in to a map. If `keys` and `values`
are not the same length, Zip will use the shorter of the two to determine
the length of the resulting map

```go
func ExampleZip() {
	keys := []string{"one", "two", "three"}
	values := []int{1, 2, 3}

	fmt.Println(fn.Zip(keys, values))
	// Output: map[one:1 three:3 two:2]
}
```

