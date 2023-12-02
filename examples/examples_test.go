package examples

import (
	"fmt"
	"fn"
	"math"
)

func ExampleAny() {
	negativeInts := []int{-1, -2, -3}
	anyPositive := fn.Any(negativeInts, func(item int, index int) bool {
		return item > 0
	})

	fmt.Println(anyPositive)
	// Output: false
}

func ExampleAnyNonZero() {
	nonZeroStrings := []string{"one", "", "", "four"}
	fmt.Println(fn.AnyNonZero(nonZeroStrings))

	zeroStrings := []string{"", ""}
	fmt.Println(fn.AnyNonZero(zeroStrings))
	// Output:
	// true
	// false
}

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

func ExampleMap() {
	ints := []int{1, 2, 3}
	doubled := fn.Map(ints, func(item int, index int) int {
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

	value := fn.Must(canFail(false))

	fmt.Println(value)
	// Output: success
}

func ExampleMax() {
	fmt.Println(fn.Max(1, 2, 100, -1))
	// Output: 100
}

func ExampleMin() {
	fmt.Println(fn.Min(-100, -300, 3, 100))
	// Output: -300
}

func ExampleReduce() {
	combined := fn.Reduce([]string{"a", "b", "c"}, func(result, item string, index int) string {
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

	evenNumbers := fn.Select(allInts, isEven)

	fmt.Println(evenNumbers)
	// Output: [2 4 6]
}

func ExampleReject() {
	allInts := []int{1, 2, 3, 4, 5, 6}

	isEven := func(number int, index int) bool {
		return number%2 == 0
	}

	oddNumbers := fn.Reject(allInts, isEven)

	fmt.Println(oddNumbers)
	// Output: [1 3 5]
}

func ExampleSum() {
	fmt.Println(fn.Sum([]int{1, 2, 3}))
	fmt.Println(fn.Sum([]float64{math.Pi, math.Pi}))
	// Output:
	// 6
	// 6.283185307179586
}

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

func ExampleZip() {
	keys := []string{"one", "two", "three"}
	values := []int{1, 2, 3}

	fmt.Println(fn.Zip(keys, values))
	// Output: map[one:1 three:3 two:2]
}
