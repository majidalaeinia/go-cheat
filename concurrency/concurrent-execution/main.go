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
	//initializing channels
	a := make(chan bool)
	b := make(chan bool)
	c := make(chan bool)
	d := make(chan bool)
	e := make(chan bool)

	//pretending that there are 5 validations
	go validation1(a)
	go validation2(b)
	go validation3(c)
	go validation4(d)
	go validation5(e)

	//receiving from each channel
	<-a
	<-b
	<-c
	<-d
	<-e
	fmt.Println("end validations")
}

func validation1(a chan bool) {
	fmt.Println("start validation: 1")
	time.Sleep(1 * time.Second)
	fmt.Println("end validation: 1")
	a <- true
}

func validation2(b chan bool) {
	fmt.Println("start validation: 2")
	time.Sleep(1 * time.Second)
	fmt.Println("end validation: 2")
	b <- true
}

func validation3(c chan bool) {
	fmt.Println("start validation: 3")
	time.Sleep(1 * time.Second)
	fmt.Println("end validation: 3")
	c <- true
}

func validation4(d chan bool) {
	fmt.Println("start validation: 4")
	time.Sleep(1 * time.Second)
	fmt.Println("end validation: 4")
	d <- true
}

func validation5(e chan bool) {
	fmt.Println("start validation: 5")
	time.Sleep(1 * time.Second)
	fmt.Println("end validation: 5")
	e <- true
}
