// Psáno v jazyce golang
package main

import "fmt"

type day struct {
	start int
	cost  int
	dur   int
}

func main() {
	// vstup dat
	var N, D int
	fmt.Scanf("%d%d", &N, &D)
	days := make([]day, D)
	for i := 0; i < D; i++ {
		var d, c, t int
		fmt.Scanf("%d%d%d", &d, &c, &t)
		days[i].start = d
		days[i].cost = c
		days[i].dur = t
		var test bool = true
		var tmp day
		for j := 0; j <= i; j++ { // insert sort
			if d < days[j].start && test {
				tmp = days[j]
				days[j] = days[i]
				test = false
			} else if !test {
				swp := tmp
				tmp = days[j]
				days[j] = swp
			}
		}
	}
	// vlastní počítání
	var byCost, byDur []day
	var cost int
	byCost = make([]day, 0)
	byDur = make([]day, 0)
	for i := 1; i <= N; i++ { // proiteruje kazdy den
		if len(days) > 0 && days[0].start == i { // přidání lahviček
			days[0].dur += (days[0].start + 1)
			byDur = append(byDur, days[0]) // přidání dle výdrže
			for j := len(byDur) - 1; j > 0; j-- {
				if byDur[j-1].dur > byDur[j].dur { // probublá lahvičku na správné místo
					swp := byDur[j-1]
					byDur[j-1] = byDur[j]
					byDur[j] = swp
				}
			}
			byCost = append(byCost, days[0]) // přidání dle ceny
			for j := len(byCost) - 1; j > 0; j-- {
				if byCost[j-1].cost > byCost[j].cost { // probublá lahvičku na správné místo
					swp := byCost[j-1]
					byCost[j-1] = byCost[j]
					byCost[j] = swp
				}
			}
			days = days[1:]
		}
		for len(byDur) > 0 && byDur[0].dur == i { // odstraneni proslych lahvicek
			var test bool = false
			for j := 0; j < len(byCost)-1; j++ {
				if len(byCost) > 0 && (byCost[j] == byDur[0] || test) { // probublání na konec, cena
					swp := byCost[j]
					byCost[j] = byCost[j+1]
					byCost[j+1] = swp
					test = true
				}
			}
			byCost = byCost[:len(byCost)-1] // odstranění koncem dle ceny
			byDur = byDur[1:] // odtranění prvního elementu, dle výdrže
		}
		if len(byCost) == 0 {
			fmt.Printf("Experiment konci dnem %d\n", i)
			return
		} else {
			cost += byCost[0].cost
		}
	}
	fmt.Printf("%d\n", cost)
}
