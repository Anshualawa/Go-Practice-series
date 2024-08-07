package main

import (
	"fmt"
)

type PersonalInfo struct {
	Name       string
	Age        int
	Height     float32
	Weight     float32
	BloodGroup string
}

var bio PersonalInfo = PersonalInfo{
	Name:       "Alawa",
	Age:        20,
	Height:     6.8,
	Weight:     78.9,
	BloodGroup: "O+",
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Printf("%+v\n", bio)
}
