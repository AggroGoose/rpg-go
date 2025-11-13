package monster

import "math/rand"

type Monster struct {
	name          string
	description   string
	aggressive    bool
	healthMax     int
	healthCurrent int
	atkDice       int
	atkSides      int
	armorClass    int
	hit           int
	exp           int
}

func CreateMonster() Monster {
	health := rand.Intn(6) + rand.Intn(6)
	return Monster{
		name:        "Goblin",
		description: "Cheeky little bugger ain't he?",
		aggressive:  false,
		healthMax:   health,
		healthCurrent: health,
		atkDice: 1,
		atkSides: 6,
		armorClass: 13,
		hit: 2,
		exp: 50,
	}
}