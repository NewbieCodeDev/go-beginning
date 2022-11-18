package main

// import one lib
import (
	"fmt"

)

var name , friendName string = "New","First"

func main() {
	numberValue := 5
	fmt.Println("Hello",name+" and " +friendName)
	fmt.Println("Number = ", numberValue)
}

// run in terminal --> go run $fileName
// go mod init $yourpath(recommand github.com/$nameOfYourGithub/$anyname) build module manage go (go.mod file). I think it same node_module file

// go code snippet --> ctrl shift p --> insert code

// go build $fileName --> file.exe can run in any pc
// go get = npm install , dogo = nodemon


// variable type in go --> Number String Boolean Array Slice Struct Pointer Function Interface Map Channel