package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("product =", product)

	//add
	product["Macbook"] = 50000
	product["IPhone"] = 30000
    fmt.Println("product =", product)

	//delete
	delete(product,"Macbook")
    fmt.Println("product =", product)

	//update
	product["IPhone"] = 25000
	fmt.Println("product =", product)

	courseName := map[string]string{"101":"Java","102":"Python"}
	fmt.Println("product =", courseName)

	//เก็บในรูปแบบ key value --> map[101:Java 102:Python]
}
