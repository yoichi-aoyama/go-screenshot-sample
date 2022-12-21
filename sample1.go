package main

import (
	"fmt"
	"time"
)

func main() {
	day := time.Now()
	const layout = "2006-01-02"
	fmt.Println(day.Format(layout)) // => "2015-05-05"
	const layout2 = "2006_01_02_15_04_05"
	fmt.Println(day.Format(layout2)) // => "2015-05-05"
}
