package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Dice represents e.g. a d4 or a -d20
type Dice struct {
	sides int
	neg   bool
}

func (d Dice) roll() int {
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Sides: %v\n", d.sides)
	res := rand.Intn(d.sides) + 1
	fmt.Printf("Result: %v\n", res)

	if d.neg == true {
		return -res
	}
	return res
}
