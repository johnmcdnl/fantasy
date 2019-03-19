package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/johnmcdnl/fantasy"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var l = fantasy.NewLeague(10)
	l.Draft()

	// for _, t := range l.Team {
	// 	fmt.Printf("%s %.2f \t %2d %2d %2d %2d %2d \n", t.Name, t.Total, t.Q1.PositionRank, t.R1.PositionRank, t.R2.PositionRank, t.W1.PositionRank, t.W2.PositionRank)
	// }

	var qb1Pos []int
	var rb1Pos []int
	var rb2Pos []int
	var wr1Pos []int
	var wr2Pos []int
	var te1Pos []int
	var kicPos []int
	var defPos []int

	for season := 1; season <= 100000; season++ {
		var l = fantasy.NewLeague(10)
		l.Draft()
		t := l.T[0]
		fmt.Printf("%2d \t %s %.2f \t %2d %2d %2d %2d %2d %2d %2d %2d \n", season, t.Name, t.Total, t.QB1.PositionRank, t.RB1.PositionRank, t.RB2.PositionRank, t.WR1.PositionRank, t.WR2.PositionRank, t.TE1.PositionRank, t.Kicker.PositionRank, t.Defence.PositionRank)

		qb1Pos = append(qb1Pos, t.QB1.PositionRank)
		rb1Pos = append(rb1Pos, t.RB1.PositionRank)
		rb2Pos = append(rb2Pos, t.RB2.PositionRank)
		wr1Pos = append(wr1Pos, t.WR1.PositionRank)
		wr2Pos = append(wr2Pos, t.WR2.PositionRank)
		te1Pos = append(te1Pos, t.TE1.PositionRank)
		kicPos = append(kicPos, t.Kicker.PositionRank)
		defPos = append(defPos, t.Defence.PositionRank)
	}

	fmt.Printf("qb1 %2d %2.0f%s \n", mean(qb1Pos), float64(mean(qb1Pos))/float64(10)*100, "%")
	fmt.Printf("rb1 %2d %2.0f%s \n", mean(rb1Pos), float64(mean(rb1Pos))/float64(20)*100, "%")
	fmt.Printf("rb2 %2d %2.0f%s \n", mean(rb2Pos), float64(mean(rb2Pos))/float64(20)*100, "%")
	fmt.Printf("wr1 %2d %2.0f%s \n", mean(wr1Pos), float64(mean(wr1Pos))/float64(20)*100, "%")
	fmt.Printf("wr2 %2d %2.0f%s \n", mean(wr2Pos), float64(mean(wr2Pos))/float64(20)*100, "%")
	fmt.Printf("te1 %2d %2.0f%s \n", mean(te1Pos), float64(mean(te1Pos))/float64(10)*100, "%")
	fmt.Printf("kic %2d %2.0f%s \n", mean(kicPos), float64(mean(kicPos))/float64(10)*100, "%")
	fmt.Printf("def %2d %2.0f%s \n", mean(defPos), float64(mean(defPos))/float64(10)*100, "%")
}

func mean(values []int) int {
	var sum int

	for _, v := range values {
		sum += v
	}

	return sum / (len(values))
}
