package main

import "fmt"

type animal interface {
	run() string
}

type bird struct{}
type dog struct{}
type fish struct{}

func main() {
	b := bird{}
	d := dog{}
	f := fish{}

	runAnimal(b)
	runAnimal(d)
	runAnimal(f)
}

func (bird) run() string {
	return "I'm a bird, I fly."
}

func (dog) run() string {
	return "I'm a dog, I walk."
}

func (fish) run() string {
	return "I'm a fish, I swim."
}

func runAnimal(a animal) {
	fmt.Println(a.run())
}
