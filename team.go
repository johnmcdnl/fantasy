package fantasy

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type Team struct {
	Name    string
	ID      int
	QB1     *QuarterBack
	RB1     *RunningBack
	RB2     *RunningBack
	WR1     *WideReceiver
	WR2     *WideReceiver
	TE1     *TightEnd
	Kicker  *Kicker
	Defence *Defence
	Total   float64
}

func (t *Team) String() string {
	j, _ := json.Marshal(t)
	return string(j)
}

func NewT(n int) *Team {
	var t Team

	t.Name = fmt.Sprintf("team #%02d", n)
	t.ID = n
	return &t
}

func (t *Team) findRandomEmptyPos() Position {

	var needs []Position

	studPositionsFilled := t.RB1 != nil && t.RB2 != nil && t.WR1 != nil && t.WR2 != nil
	//coreCoreFilled := t.RB1 != nil && t.RB2 != nil && t.WR1 != nil && t.WR2 != nil && t.QB1 != nil
	coreFilled := t.RB1 != nil && t.RB2 != nil && t.WR1 != nil && t.WR2 != nil && t.QB1 != nil && t.TE1 != nil

	if studPositionsFilled && t.QB1 == nil {
		needs = append(needs, PosQB1)
	}
	if t.RB1 == nil {
		needs = append(needs, PosRB1)
	}
	if t.RB1 != nil && t.RB2 == nil {
		needs = append(needs, PosRB2)
	}
	if t.WR1 == nil {
		needs = append(needs, PosWR1)
	}
	if t.WR1 != nil && t.WR2 == nil {
		needs = append(needs, PosWR2)
	}
	if t.TE1 == nil {
		needs = append(needs, PosTE)
	}
	if coreFilled && t.Kicker == nil {
		needs = append(needs, PosK)
	}
	if coreFilled && t.Defence == nil {
		needs = append(needs, PosDEF)
	}

	return needs[rand.Intn(len(needs))]
}
