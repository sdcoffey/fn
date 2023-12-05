package enum

import "fmt"

func ExampleEnumeratorFromFunc() {
	e := fromFunc(0, func(lastValue int) (int, bool) {
		newValue := lastValue + 1
		return newValue, newValue < 10
	})

	fmt.Println(e.ReadAll())
	// Output:
	// [0 1 2 3 4 5 6 7 8 9]
}

func ExampleSelect() {
	seq(0, 5, 1).selectEnum(func(item int) bool {
		return item%2 == 0
	}).Each(func(item int) {
		fmt.Println(item)
	})

	// Output:
	// 0
	// 2
	// 4
}
