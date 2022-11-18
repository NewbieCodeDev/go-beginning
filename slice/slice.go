package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Java", "Python"}
	fmt.Println(courseName)
	courseName = append(courseName, "C", "C++")
	fmt.Println(courseName)

	courseWeb := courseName[1:4]
	fmt.Println(courseWeb)

	// ไม่จำเป็นต้องระบุจำนวนขนาดก่อน
}
