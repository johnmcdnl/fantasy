package fantasy

type Player struct {
	Position     string
	PositionRank int
	Points       float64
}

type QuarterBack struct{ *Player }
type RunningBack struct{ *Player }
type WideReceiver struct{ *Player }
type TightEnd struct{ *Player }
type Kicker struct{ *Player }
type Defence struct{ *Player }

func NextQB(p *Projection) *QuarterBack  { return &QuarterBack{nextPlayer(p, "QB"),} }
func NextRB(p *Projection) *RunningBack  { return &RunningBack{nextPlayer(p, "RB"),} }
func NextWR(p *Projection) *WideReceiver { return &WideReceiver{nextPlayer(p, "WR"),} }
func NextTE(p *Projection) *TightEnd     { return &TightEnd{nextPlayer(p, "TE"),} }
func NextKicker(p *Projection) *Kicker   { return &Kicker{nextPlayer(p, "KR"),} }
func NextDefence(p *Projection) *Defence { return &Defence{nextPlayer(p, "DF"),} }

func nextPlayer(p *Projection, pos string) *Player {
	defer p.CalculateNextFunc()
	return &Player{
		Position:     pos,
		PositionRank: p.NextRank,
		Points:       p.NextPoints,
	}
}

type Position int

const (
	PosQB1 = iota
	PosRB1
	PosRB2
	PosWR1
	PosWR2
	PosTE
	PosK
	PosDEF
)
