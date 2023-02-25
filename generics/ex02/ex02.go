package main

import "fmt"

type Service struct{}

type Car struct {
	Speed int
}
type Motocycle struct {
	Speed float64
}

type Running interface {
	*Car | *Motocycle
}

func Printer[T Running](s T) {
	fmt.Println(s)
}

func main() {
	Printer(&Car{})
	Printer(&Motocycle{})
}
