package main

import "fmt"

func plus(value, value1 int) int {
	return value + value1

}

func main() {
	result := plus(1, 2)
	fmt.Println("result =", result)
}
