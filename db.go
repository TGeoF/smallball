package main

func loadLeague() League {
	p1 := Player{name: "p1", BT: 27, OBT: 32, pitchDice: Dice{sides: 4, neg: false}}
	p2 := Player{name: "p2", BT: 27, OBT: 32, pitchDice: Dice{sides: 4, neg: false}}
	p3 := Player{name: "p3", BT: 27, OBT: 32, pitchDice: Dice{sides: 4, neg: false}}
	p4 := Player{name: "p4", BT: 27, OBT: 32, pitchDice: Dice{sides: 4, neg: false}}
	p5 := Player{name: "p5", BT: 27, OBT: 32, pitchDice: Dice{sides: 4, neg: false}}
	p6 := Player{name: "p6", BT: 27, OBT: 32, pitchDice: Dice{sides: 4, neg: false}}
	p7 := Player{name: "p7", BT: 27, OBT: 32, pitchDice: Dice{sides: 4, neg: false}}
	p8 := Player{name: "p8", BT: 27, OBT: 32, pitchDice: Dice{sides: 4, neg: false}}
	p9 := Player{name: "p9", BT: 27, OBT: 32, pitchDice: Dice{sides: 4, neg: false}}

	oak := Team{nickname: "Athletics",
		roster:    []*Player{&p1, &p2, &p3, &p4, &p5, &p6, &p7, &p8, &p9},
		lineup:    [9]*Player{&p1, &p2, &p3, &p4, &p5, &p6, &p7, &p8, &p9},
		positions: [9]*Player{&p1, &p2, &p3, &p4, &p5, &p6, &p7, &p8, &p9}}
	nym := Team{nickname: "Mets",
		roster:    []*Player{&p1, &p2, &p3, &p4, &p5, &p6, &p7, &p8, &p9},
		lineup:    [9]*Player{&p1, &p2, &p3, &p4, &p5, &p6, &p7, &p8, &p9},
		positions: [9]*Player{&p1, &p2, &p3, &p4, &p5, &p6, &p7, &p8, &p9}}

	return League{teams: map[string]Team{"oak": oak, "nym": nym}}
}
