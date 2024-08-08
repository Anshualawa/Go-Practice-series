package main

import "fmt"

type Comm interface {
	finul() int
}

type Sqr struct {
	Num int
}

func (s Sqr) finul() int {
	return s.Num * s.Num
}

type Cub struct {
	Num int
}

func (c Cub) finul() int {
	return c.Num * c.Num * c.Num
}

func Commfun(c Comm) {
	fmt.Println(c.finul())
}

func main() {
	sq := Sqr{5}
	cub := Cub{3}
	Commfun(sq)
	Commfun(cub)

}
