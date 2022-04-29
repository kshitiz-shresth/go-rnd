package main

import (
	"fmt"
	"time"
)

func main() {
	var currentTime = time.Now()
	fmt.Println(currentTime.Format("2006-1-02")) //got surprised because it gave output of today's date
}
