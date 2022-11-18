package main

import "fmt"

func changeValueByPoint(n *int) {
	*n = 0
}
func main() {
	i := 5

	fmt.Println("before change i =", i)
	
	changeValueByPoint(&i)
	fmt.Println("after changed i =", i)

	// & คือ เอาที่อยู่ของตัวแปรนั้นๆ , * เอาค่าจากที่อยู่
}
