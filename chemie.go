package main

import "fmt"

type day struct {
	start int
	cost  int
	dur   int
}

func main() {
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
			if d < days[j].dur && test {
				tmp = days[j]
				days[j] = days[i]
				test = false
			}
			if !test {
				swp := tmp
				tmp = days[j]
				days[j] = swp
			}
		}
	}
}
