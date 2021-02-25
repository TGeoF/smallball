package main

import "fmt"

type handler interface {
	handle() Event
}

// PaHandler can apply the rules to a game, effectively progressing the game by one plate appearance
type PaHandler struct {
}

func (p PaHandler) handle(g *Game, r *Rules) (PaEvent, error) {
	e := PaEvent{}
	d100 := Dice{sides: 100}
	var batter *Player
	var pitcher *Player

	fmt.Printf("\nInning: %v%v\nOuts: %v\n", g.halfInning, g.inning+1, g.outs)

	batter, pitcher = g.getMatchup()

	fmt.Printf("Pitching: %v\n", pitcher.name)
	fmt.Printf("Batting: %v\n", batter.name)

	result := d100.roll() - pitcher.pitchDice.roll()

	fmt.Printf("Dice Result: %v\n", result)

	if result <= batter.BT {
		e = r.paEvents["single"]
	} else if result <= batter.OBT {
		e = r.paEvents["walk"]
	} else {
		e = r.paEvents["strikeout"]
	}

	fmt.Println(e.log())

	if e.batterOut {
		g.outs++
	}

	if e.code == "BB" || e.code == "HBP" {
		g.grantedFirst(batter)
	}

	if e.hit {
		g.hit(e.bases)
	}

	if g.outs == 3 {
		g.nextInning()
	}

	return e, nil
}
