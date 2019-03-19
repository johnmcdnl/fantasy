package fantasy

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
)

type League struct {
	T            []*Team
	qbProjection *Projection
	rbProjection *Projection
	wrProjection *Projection
	teProjection *Projection
	krProjection *Projection
	dfProjection *Projection
}

func NewLeague(n int) League {

	var l League

	l.qbProjection = newProjection(400, 2.1)
	l.rbProjection = newProjection(300, 1.1)
	l.wrProjection = newProjection(320, 1.1)
	l.teProjection = newProjection(200, 1)
	l.krProjection = newProjection(160, 1)
	l.dfProjection = newProjection(160, 0.8)

	for i := 1; i <= n; i++ {
		l.T = append(l.T, NewT(i))
	}

	return l
}

func (l *League) String() string {
	j, _ := json.Marshal(l)
	return string(j)
}

func (l *League) Draft() {
	var smartTeam = "team JMD"
	l.T[rand.Intn(len(l.T))].Name = smartTeam

	for round := 1; round <= 8; round++ {
		for _, t := range l.T {
			if t.Name == smartTeam+"NOT" {
				switch round {
				case 1:
					t.WR1 = l.NextWR()
				case 2:
					t.RB1 = l.NextRB()
				case 3:
					t.WR2 = l.NextWR()
				case 4:
					t.RB2 = l.NextRB()
				case 5:
					t.TE1 = l.NextTE()
				case 6:
					t.Defence = l.NextDef()
				case 7:
					t.QB1 = l.NextQB()
				case 8:
					t.Kicker = l.NextK()
				}
				continue
			}
			switch t.findRandomEmptyPos() {
			default:
				panic("Unexpected")
			case PosQB1:
				t.QB1 = l.NextQB()
			case PosRB1:
				t.RB1 = l.NextRB()
			case PosRB2:
				t.RB2 = l.NextRB()
			case PosWR1:
				t.WR1 = l.NextWR()
			case PosWR2:
				t.WR2 = l.NextWR()
			case PosTE:
				t.TE1 = l.NextTE()
			case PosK:
				t.Kicker = l.NextK()
			case PosDEF:
				t.Defence = l.NextDef()
			}
		}
		// Snake draft
		for i := len(l.T)/2 - 1; i >= 0; i-- {
			opp := len(l.T) - 1 - i
			l.T[i], l.T[opp] = l.T[opp], l.T[i]
		}
	}

	for _, t := range l.T {
		t.Total = t.QB1.Points + t.RB1.Points + t.RB2.Points + t.WR1.Points + t.WR2.Points + t.TE1.Points + t.Kicker.Points + t.Defence.Points
	}

	sort.Slice(l.T, func(i, j int) bool {
		return l.T[i].Total > l.T[j].Total
	})

	totalCount++

	if l.T[0].Name == smartTeam {
		winCount++
		fmt.Println(winCount, "/", totalCount, float64(winCount)/float64(totalCount)*100)
	}
}

var winCount int
var totalCount int

func (l *League) NextQB() *QuarterBack {
	return NextQB(l.qbProjection)
}

func (l *League) NextRB() *RunningBack {
	return NextRB(l.rbProjection)
}

func (l *League) NextWR() *WideReceiver {
	return NextWR(l.wrProjection)
}

func (l *League) NextTE() *TightEnd {
	return NextTE(l.teProjection)
}

func (l *League) NextK() *Kicker {
	return NextKicker(l.krProjection)
}

func (l *League) NextDef() *Defence {
	return NextDefence(l.dfProjection)
}
