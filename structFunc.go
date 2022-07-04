package main

import "fmt"

type MyStruct struct {
	Name string
}

func (this *MyStruct) fnInsideStruct() {
	fmt.Println("Hello Sir", this.Name)
}

type Shop struct {
	TotalProfit int32
	TotalLoss   int32
}

func (this *Shop) getGrossProfit(myStruct MyStruct) int32 {
	myStruct.Name = "Ram"
	this.TotalProfit = 100
	return this.TotalProfit - this.TotalLoss
}

func main() {
	var tempStr MyStruct
	tempStr.Name = "Gopal"

	var coldStore = Shop{300, 300}
	fmt.Println(coldStore.getGrossProfit(tempStr))
	fmt.Println(coldStore.TotalProfit)
	tempStr.fnInsideStruct()

}
