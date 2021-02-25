package main

// Rules implements a ruleset containing events and tables
type Rules struct {
	paEvents map[string]PaEvent
}

// Event represents a possible outcome of a decision or situation
type Event interface {
	log() string
	getCode() string
}

// PaEvent represents the outcome of a plate appearence
type PaEvent struct {
	representation string
	code           string
	atBat          bool
	onBase         bool
	hit            bool
	trueOutcome    bool
	bases          int
	batterOut      bool
	liveBall       bool
}

func (p *PaEvent) log() string {
	return p.code
}

func (p *PaEvent) getCode() string {
	return p.code
}

func (r *Rules) load() {
	walk := PaEvent{code: "BB",
		atBat:       false,
		onBase:      true,
		hit:         false,
		trueOutcome: true,
		bases:       1,
		batterOut:   false,
		liveBall:    false}
	hitbypitch := PaEvent{code: "HBP",
		atBat:       false,
		onBase:      true,
		hit:         false,
		trueOutcome: true,
		bases:       1,
		batterOut:   false,
		liveBall:    false}
	strikeout := PaEvent{code: "K",
		atBat:       true,
		onBase:      false,
		hit:         false,
		trueOutcome: true,
		bases:       0,
		batterOut:   true,
		liveBall:    false}
	single := PaEvent{code: "1B",
		atBat:       true,
		onBase:      true,
		hit:         true,
		trueOutcome: false,
		bases:       1,
		batterOut:   false,
		liveBall:    true}
	double := PaEvent{code: "2B",
		atBat:       true,
		onBase:      true,
		hit:         true,
		trueOutcome: false,
		bases:       2,
		batterOut:   false,
		liveBall:    true}
	triple := PaEvent{code: "3B",
		atBat:       true,
		onBase:      true,
		hit:         true,
		trueOutcome: false,
		bases:       3,
		batterOut:   false,
		liveBall:    true}
	homerun := PaEvent{code: "HR",
		atBat:       true,
		onBase:      true,
		hit:         true,
		trueOutcome: true,
		bases:       4,
		batterOut:   false,
		liveBall:    false}

	r.paEvents = map[string]PaEvent{"walk": walk,
		"hitbypitch": hitbypitch,
		"single":     single,
		"double":     double,
		"triple":     triple,
		"homerun":    homerun,
		"strikeout":  strikeout}
}
