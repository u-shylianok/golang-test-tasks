package main

import "fmt"

// Rocket with only one Engine
type Rocket struct {
	Name   string
	Year   int
	Price  int // not decimal to simplify
	Weight int // kg
	Engine
}

func (r Rocket) String() string {
	return fmt.Sprintf("Rocket name: %s, %d. -- Price: %d, Weight: %d. -- Engine: [%v]", r.Name, r.Year, r.Price, r.Weight, r.Engine)
}

// My own interpretation of Rocket Move
func (r Rocket) Move() {
	fmt.Println("Rocket: *", r.Name, "* moves with speed:", r.MagicPower/GetAllWeight(r), "m/s")
}

// Returns rocket + engine weight
func GetAllWeight(r Rocket) int {
	return r.Engine.Weight + r.Weight
}

// Create some (5) rockets with constant values
func CreateSomeRockets() [5]Rocket {
	trueMagicEngine := Engine{
		Name:       "True Magic",
		MagicPower: 999999,
		Weight:     50,
	}
	fakeMagicEngine := Engine{
		Name:       "Fake Magic",
		MagicPower: 30000,
		Weight:     5,
	}

	rockets := [5]Rocket{
		{
			Name:   "Rr1",
			Year:   2001,
			Price:  5000,
			Weight: 100,
			Engine: fakeMagicEngine,
		},
		{
			Name:   "Rsuper1",
			Year:   2002,
			Price:  250500,
			Weight: 1400,
			Engine: trueMagicEngine,
		},
		{
			Name:   "Rsuper2",
			Year:   2003,
			Price:  110000,
			Weight: 1350,
			Engine: trueMagicEngine,
		},
		{
			Name:   "Rr2",
			Year:   2005,
			Price:  100100,
			Weight: 15,
			Engine: fakeMagicEngine,
		},
	}
	// after 2010
	trueMagicEngine.MagicPower *= 3
	rockets[4] = Rocket{
		Name:   "Rsuper3",
		Year:   2012,
		Price:  500500,
		Weight: 1400,
		Engine: trueMagicEngine,
	}
	return rockets
}
