package main

import (
	"fmt"
)

// Simple-magic rocket engine
type Engine struct {
	Name       string
	MagicPower int // can move a rocket at a constant speed by the formula: MP/AllWeight
	Price      int // not decimal to simplify
	Weight     int // kg
}

// Rocket with only one Engine
type Rocket struct {
	Name   string
	Year   int
	Price  int // not decimal to simplify
	Weight int // kg
	Engine
}

// My own interpretation of Rocket Move
func (r Rocket) Move() {
	fmt.Println("Rocket: *", r.Name, "* moves with speed:", r.MagicPower/GetAllWeight(r), "m/s")
}

// Returns rocket + engine weight
func GetAllWeight(r Rocket) int {
	return r.Engine.Weight + r.Weight
}

func main() {
	fmt.Println("--- task01 ---")
	rocket := Rocket{
		Name:   "Big Boy",
		Year:   2000,
		Price:  3000000,
		Weight: 1500,
		Engine: Engine{
			Name:       "True Magic",
			MagicPower: 999999,
			Weight:     50,
		}}

	fmt.Println("AllWeight:", GetAllWeight(rocket), "kg")
	rocket.Move()
}
