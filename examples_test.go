package fn

import (
	"errors"
	"fmt"
	"math"
)

func ExampleAny() {
	negativeInts := []int{-1, -2, -3}
	anyPositive := Any(negativeInts, func(item int, index int) bool {
		return item > 0
	})

	fmt.Println(anyPositive)
	// Output: false
}

func ExampleAnyNonZero() {
	nonZeroStrings := []string{"one", "", "", "four"}
	fmt.Println(AnyNonZero(nonZeroStrings))

	zeroStrings := []string{"", ""}
	fmt.Println(AnyNonZero(zeroStrings))
	// Output:
	// true
	// false
}

func ExampleAll() {
	positiveInts := []int{1, 2, 3}
	allPositive := All(positiveInts, func(item int, index int) bool {
		return item > 0
	})

	fmt.Println(allPositive)
	// Output: true
}

func ExampleAllNonZero() {
	nonZeroInts := []string{"one", "two", "three"}
	fmt.Println(AllNonZero(nonZeroInts))

	zeroInts := []string{"", "", ""}
	fmt.Println(AllNonZero(zeroInts))
	// Output:
	// true
	// false
}

func ExampleChunk() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(Chunk(ints, 3))
	// Output: [[1 2 3] [4 5 6] [7 8 9]]
}

func ExampleChunkWhile() {
	ints := []int{1, 2, 4, 5, 7}

	fmt.Println(ChunkWhile(ints, func(eltBefore, eltAfter int) bool {
		return eltBefore+1 == eltAfter
	}))
	// Output: [[1 2] [4 5] [7]]
}

func ExampleCompactZero() {
	ints := []int{0, 1, 2, 3, 0, 5}

	fmt.Println(CompactZero(ints))
	// Output: [1 2 3 5]
}

func ExampleCompactNil() {
	errs := []error{nil, errors.New("example-error 1"), nil, errors.New("example-error 2")}

	fmt.Println(CompactNil(errs))
	// Output: [example-error 1 example-error 2]
}

func ExampleEach() {
	Each([]string{"a", "b", "c"}, func(item string, index int) {
		fmt.Println(item, index)
	})
	// Output:
	// a 0
	// b 1
	// c 2
}

func ExampleFirst() {
	sequence := []int{-1, 0, 1, 2}
	firstPositive, index := First(sequence, func(item int, index int) bool {
		return item > 0
	})

	ten, tenIndex := First(sequence, func(item int, index int) bool {
		return item == 10
	})

	fmt.Printf("value %d at index %d\n", firstPositive, index)
	fmt.Printf("value %d at index %d\n", ten, tenIndex)
	// Output:
	// value 1 at index 2
	// value 0 at index -1
}

func ExampleFlatten() {
	items := [][]string{{"one"}, {"two", "three"}, {"four"}}

	fmt.Println(Flatten(items))
	// Output: [one two three four]
}

func ExampleMap() {
	ints := []int{1, 2, 3}
	doubled := Map(ints, func(item int, index int) int {
		return item * 2
	})

	fmt.Println(doubled)
	// Output: [2 4 6]
}

func ExampleMust() {
	canFail := func(shouldFail bool) (string, error) {
		if shouldFail {
			return "", fmt.Errorf("error")
		}
		return "success", nil
	}

	value := Must(canFail(false))

	fmt.Println(value)
	// Output: success
}

func ExampleMax() {
	fmt.Println(Max(1, 2, 100, -1))
	// Output: 100
}

func ExampleMin() {
	fmt.Println(Min(-100, -300, 3, 100))
	// Output: -300
}

func ExamplePartition() {
	values := Seq(0, 10, 1)
	evens, odds := Partition(values, func(item, index int) bool {
		return item%2 == 0
	})

	fmt.Println(evens)
	fmt.Println(odds)
	// Output:
	// [0 2 4 6 8]
	// [1 3 5 7 9]
}

func ExampleReduce() {
	combined := Reduce([]string{"a", "b", "c"}, func(result, item string, index int) string {
		return result + item
	}, "")

	fmt.Println(combined)
	// Output: abc
}

func ExampleSelect() {
	allInts := []int{1, 2, 3, 4, 5, 6}

	isEven := func(number int, index int) bool {
		return number%2 == 0
	}

	evenNumbers := Select(allInts, isEven)

	fmt.Println(evenNumbers)
	// Output: [2 4 6]
}

func ExampleSeq() {
	seq := Seq(0, 50, 5)

	fmt.Println(seq)
	// Output: [0 5 10 15 20 25 30 35 40 45]
}

func ExampleReject() {
	allInts := []int{1, 2, 3, 4, 5, 6}

	isEven := func(number int, index int) bool {
		return number%2 == 0
	}

	oddNumbers := Reject(allInts, isEven)

	fmt.Println(oddNumbers)
	// Output: [1 3 5]
}

func ExampleSum() {
	fmt.Println(Sum([]int{1, 2, 3}))
	fmt.Println(Sum([]float64{math.Pi, math.Pi}))
	// Output:
	// 6
	// 6.283185307179586
}

func ExampleZip() {
	keys := []string{"one", "two", "three"}
	values := []int{1, 2, 3}

	fmt.Println(Zip(keys, values))
	// Output: map[one:1 three:3 two:2]
}
