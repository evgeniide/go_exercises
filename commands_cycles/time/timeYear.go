package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now() // func
	var year int = now.Year()      // method
	fmt.Println(year)
}
