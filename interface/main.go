package main

import "fmt"

// Interface
type Comm interface {
	finul() int
}
type Vehicle interface {
	Speed() int
	Fuel() string
}

// structs
type Sqr struct {
	Num int
}

type Cub struct {
	Num int
}

type Car struct{}
type Bike struct{}

// reciver functions
func (s Sqr) finul() int {
	return s.Num * s.Num
}

func (c Cub) finul() int {
	return c.Num * c.Num * c.Num
}

func (c Car) Speed() int {
	return 120
}
func (c Car) Fuel() string {
	return "CNG and Petrol"
}
func (c Bike) Speed() int {
	return 80
}
func (c Bike) Fuel() string {
	return "Petrol"
}

// function
func Commfun(c Comm) {
	fmt.Println(c.finul())
}

func main() {
	sq := Sqr{5}
	cub := Cub{3}
	Commfun(sq)
	Commfun(cub)

	car := Car{}
	bike := Bike{}

	// an array of the Vehicle interface type
	vehicles := []Vehicle{car, bike}

	for _, vehicl := range vehicles {
		fmt.Printf("Speed: %d km/h, Fuel: %s\n", vehicl.Speed(), vehicl.Fuel())
	}

}
