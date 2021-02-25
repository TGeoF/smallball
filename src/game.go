package main

// Game represents a game of baseball between two teams and its current gamestate
type Game struct {
	Home       Team
	Away       Team
	inning     int
	halfInning string
	nextHome   int
	nextAway   int
	runsHome   [9]int
	runsAway   [9]int
	basepath   [3]*Player
	outs       int
}

func (g *Game) init() {
	g.halfInning = "t"
	g.runsAway = [9]int{}
	g.runsHome = [9]int{}
	g.nextAway = 0
	g.nextHome = 0
	g.basepath = [3]*Player{nil, nil, nil}
	g.outs = 0
}

func (g *Game) nextInning() error {
	if g.halfInning == "t" {
		g.halfInning = "b"
	} else if g.halfInning == "b" {
		g.halfInning = "t"
		g.inning++
	}
	g.outs = 0
	g.basepath = [3]*Player{nil, nil, nil}
	return nil
}

func (g *Game) getMatchup() (*Player, *Player) {
	var batter, pitcher *Player
	if g.halfInning == "t" {
		batter = g.Away.lineup[g.nextAway]
		pitcher = g.Home.positions[0]
		if g.nextAway == 8 {
			g.nextAway = 0
		} else {
			g.nextAway++
		}
	} else if g.halfInning == "b" {
		batter = g.Home.lineup[g.nextHome]
		pitcher = g.Away.positions[0]
		if g.nextHome == 8 {
			g.nextHome = 0
		} else {
			g.nextHome++
		}
	}
	return batter, pitcher
}

func (g *Game) grantedFirst(batter *Player) error {
	if g.basepath[0] == nil {
		g.basepath[0] = batter
	} else if g.getBasesLoaded() {
		g.scoreRuns(1)
		g.basepath[2] = g.basepath[1]
		g.basepath[1] = g.basepath[0]
		g.basepath[0] = batter
	} else if g.basepath[1] == nil {
		g.basepath[1] = g.basepath[0]
		g.basepath[0] = batter
	} else {
		g.basepath[2] = g.basepath[1]
		g.basepath[1] = g.basepath[0]
		g.basepath[0] = batter
	}
	return nil
}

func (b *Basepath) hit(bases int) error {
	return nil
}

func (g *Game) scoreRuns(n int) {
	if g.halfInning == "t" {
		g.runsAway[g.inning] += n
	} else if g.halfInning == "b" {
		g.runsHome[g.inning] += n
	}
}

func (g *Game) getBasesLoaded() bool {
	return g.basepath[0] != nil && g.basepath[1] != nil && g.basepath[2] != nil
}
