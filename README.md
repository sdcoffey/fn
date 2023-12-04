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
* [Chunk](#Chunk)
* [ChunkWhile](#ChunkWhile)
* [CompactNil](#CompactNil)
* [CompactZero](#CompactZero)
* [Each](#Each)
* [First](#First)
* [Flatten](#Flatten)
* [Map](#Map)
* [Max](#Max)
* [Min](#Min)
* [Must](#Must)
* [Partition](#Partition)
* [Reduce](#Reduce)
* [Reject](#Reject)
* [Select](#Select)
* [Seq](#Seq)
* [Sum](#Sum)
* [Zip](#Zip)


### Any
Any returns true if any item in the slice satisfies the predicate.

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
is not the zero value for its type.

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

### Chunk
Chunk takes `[]T` and returns `[][]T`, where each subarray has a max length
of `chunkSize`. If `len(items) % chunkSize != 0`, the last slice will be
shorter than `chunkSize`.

```go
func ExampleChunk() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(fn.Chunk(ints, 3))
	// Output: [[1 2 3] [4 5 6] [7 8 9]]
}
```

### ChunkWhile
ChunkWhile takes `[]T` and returns `[][]T`, each of which is a contiguous
chunk of items for which `chunker` returns true. `chunker` is called with
the item before and the item after the current item.

```go
func ExampleChunkWhile() {
	ints := []int{1, 2, 4, 5, 7}

	fmt.Println(fn.ChunkWhile(ints, func(eltBefore, eltAfter int) bool {
		return eltBefore+1 == eltAfter
	}))
	// Output: [[1 2] [4 5] [7]]
}
```

### CompactNil
CompactNil returns a new slice with all nil items removed.

```go
func ExampleCompactNil() {
	errs := []error{nil, errors.New("example-error 1"), nil, errors.New("example-error 2")}

	fmt.Println(fn.CompactNil(errs))
	// Output: [example-error 1 example-error 2]
}
```

### CompactZero
CompactZero returns a new slice with all zero items removed.

```go
func ExampleCompactZero() {
	ints := []int{0, 1, 2, 3, 0, 5}

	fmt.Println(fn.CompactZero(ints))
	// Output: [1 2 3 5]
}
```

### Each
Each calls the given function for each item in the slice.

```go
func ExampleEach() {
	fn.Each([]string{"a", "b", "c"}, func(item string, index int) {
		fmt.Println(item, index)
	})
	// Output:
	// a 0
	// b 1
	// c 2
}
```

### First
First will return the first value that satisfies the predicate, and the
index at which the value was found (or -1 if it was not found).

```go
func ExampleFirst() {
	sequence := []int{-1, 0, 1, 2}
	firstPositive, index := fn.First(sequence, func(item int, index int) bool {
		return item > 0
	})

	ten, tenIndex := fn.First(sequence, func(item int, index int) bool {
		return item == 10
	})

	fmt.Printf("value %d at index %d\n", firstPositive, index)
	fmt.Printf("value %d at index %d\n", ten, tenIndex)
	// Output:
	// value 1 at index 2
	// value 0 at index -1
}
```

### Flatten
Flatten takes a 2D slice `[][]T` and flattens it to one dimension,
preserving the original order.

```go
func ExampleFlatten() {
	items := [][]string{{"one"}, {"two", "three"}, {"four"}}

	fmt.Println(fn.Flatten(items))
	// Output: [one two three four]
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
the set type if it has len 0.

```go
func ExampleMax() {
	fmt.Println(fn.Max(1, 2, 100, -1))
	// Output: 100
}
```

### Min
Min returns the minimum value in some orderable set, or the zero value of
the set type if it has len 0.

```go
func ExampleMin() {
	fmt.Println(fn.Min(-100, -300, 3, 100))
	// Output: -300
}
```

### Must
Must allows you to return one value from a function that would normally
return a value and an error. If the error is present, Must will panic.

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

### Partition
Partition takes `[]T` and a predicate function and returns two slices of
`T`. The first slice contains all the items for which `pred` returned true,
and the second slice contains all the items for which `pred` returned false.

```go
func ExamplePartition() {
	values := fn.Seq(0, 10, 1)
	evens, odds := fn.Partition(values, func(item, index int) bool {
		return item%2 == 0
	})

	fmt.Println(evens)
	fmt.Println(odds)
	// Output:
	// [0 2 4 6 8]
	// [1 3 5 7 9]
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
Reject reduces a slice to only values for which `rejector` returns false.

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
Select reduces a slice to only values for which `selector` returns true.

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
Seq Returns a slice of integers from `[start, end)`, incremented by `inc`.

```go
func ExampleSeq() {
	seq := fn.Seq(0, 50, 5)

	fmt.Println(seq)
	// Output: [0 5 10 15 20 25 30 35 40 45]
}
```

### Sum
Sum returns the summed value of values in the slice.

```go
func ExampleSum() {
	fmt.Println(fn.Sum([]int{1, 2, 3}))
	fmt.Println(fn.Sum([]float64{math.Pi, math.Pi}))
	// Output:
	// 6
	// 6.283185307179586
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

