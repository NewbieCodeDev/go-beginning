package main

import "fmt"

var productName [4]string

func main() {
	price := [4]float32{40000}
	productName[0] = "Macbook"
	fmt.Println(productName)
	fmt.Println(price)
}

// ต้องกำหนดขนาดอาเรย์