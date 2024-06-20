package main

import "fmt"

func hello() {
	fmt.Println("Hello brontodev")
}

func plus(value1 int, value2 int) int {
	return value1 + value2
	// result := value1 + value2
	// fmt.Println("result = ", result)
}

func plus3value(value1, value2, value3 int) int {
	return value1 + value2 + value3
}

func main() {
	hello()
	// plus(1,8)
	result := plus(1, 8)
	fmt.Println(result)

	result2 := plus3value(1, 8, 5)
	fmt.Println(result2)
}
