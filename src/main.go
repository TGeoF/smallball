package main

import "fmt"

func main() {
	l := loadLeague()
	r := Rules{}
	p := PaHandler{}
	g := Game{Home: l.teams["oak"], Away: l.teams["nym"]}

	r.load()
	g.init()

	for g.inning < 9 {
		_, err := p.handle(&g, &r)
		if err != nil {
			fmt.Println("Error")
		}
	}

	fmt.Printf("FINAL:\n%v\t\t%v\n%v\t%v", g.Away.nickname, g.runsAway,
		g.Home.nickname, g.runsHome)
}
