package main

// League represents a set of teams
type League struct {
	teams map[string]Team
}

// Team contains pointers to players as well as standings information
type Team struct {
	city      string
	nickname  string
	short     string
	roster    []*Player
	lineup    [9]*Player
	positions [9]*Player
	wins      int
	losses    int
}

// Player represents baseball players and their talent level
type Player struct {
	name      string
	BT        int
	OBT       int
	pitchDice Dice
	pos       []int
}
