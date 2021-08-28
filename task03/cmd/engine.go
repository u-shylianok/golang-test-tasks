package main

import "fmt"

// Simple-magic rocket engine
type Engine struct {
	Name       string
	MagicPower int // can move a rocket at a constant speed by the formula: MP/AllWeight
	Price      int // not decimal to simplify
	Weight     int // kg
}

func (e Engine) String() string {
	return fmt.Sprintf("Name: %s, Power: %d MP. -- Price: %d, Weight: %d", e.Name, e.MagicPower, e.Price, e.Weight)
}
