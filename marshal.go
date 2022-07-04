package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	Name    string `json:"name"`
	Contact int64  `json:"contact"`
}

func main() {
	jsonData := `{"name":"Hari","contact":9861513956}`

	t := jsonDecode(jsonData)

	fmt.Println(t)

	fmt.Println(stringify(t))

	fmt.Println(printInt(23))

}

func printInt(number int32) int32 {
	a := number
	return a
}

func jsonDecode(jsonData string) *T {
	var t *T
	t = new(T)
	err := json.Unmarshal([]byte(jsonData), t)
	if err != nil {
		t.Name = "Error"
		t.Contact = 404
	}
	return t
}

func stringify(t *T) string {
	marshalled, _ := json.Marshal(t)
	return string(marshalled)
}
