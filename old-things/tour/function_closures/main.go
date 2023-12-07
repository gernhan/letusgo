package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	sum := 0
	cnt := 0
	toReturn := func() int {
		cnt = cnt + 1
		sum += cnt
		return sum
	}
	return toReturn
}

func main() {
	pos, neg := adder(), adder()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(
	//		pos(i),
	//		neg(-2*i),
	//
	//}

	fmt.Println(
		pos(1),
		neg(-2*1),
	)

	fmt.Println(
		pos(3),
		neg(-2*4),
	)

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
