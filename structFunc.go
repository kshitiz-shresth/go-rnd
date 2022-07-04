package main

import "fmt"

type MyStruct struct {
	Name string
}

func (this *MyStruct) fnInsideStruct() {
	fmt.Println("Hello Sir", this.Name)
}

func main() {
	var tempStr = new(MyStruct)
	tempStr.Name = "Gopal"
	tempStr.fnInsideStruct()
}
