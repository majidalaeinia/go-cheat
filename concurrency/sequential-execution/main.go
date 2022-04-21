package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("start main goroutine")
	validate()
	fmt.Println("end main goroutine")
	elapsed := time.Since(start)
	fmt.Printf("program took %s", elapsed)
}

func validate() {
	fmt.Println("start validations")
	//pretending that there are 5 validations
	for i := 1; i <= 5; i++ {
		validation(i)
	}
	fmt.Println("end validations")
}

func validation(i int) {
	fmt.Println("start validation:", i)
	time.Sleep(1 * time.Second)
	fmt.Println("end validation:", i)
}
