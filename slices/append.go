package main

import "fmt"

func main() {
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
	}
}

// output

// 0       cap=1   [0]
// 1       cap=2   [0 1]
// 2       cap=4   [0 1 2]
// 3       cap=4   [0 1 2 3]
// 4       cap=8   [0 1 2 3 4]
// 5       cap=8   [0 1 2 3 4 5]
// 6       cap=8   [0 1 2 3 4 5 6]
// 7       cap=8   [0 1 2 3 4 5 6 7]
// 8       cap=16  [0 1 2 3 4 5 6 7 8]
// 9       cap=16  [0 1 2 3 4 5 6 7 8 9]

// After the 3rd iteration we can see the capacity is 4, and not 3 as we would imagine
// When a slice doesn't have enough capacity to hold more elements, Go doubles its capacity.
// It creates a new underlying array with the new capacity. You don't have to do anything for this increase in capacity to happen. Go does it automatically.
// You do need to be cautious. At some point, a slice might have way more capacity than it needs, and you'll be wasting memory.
