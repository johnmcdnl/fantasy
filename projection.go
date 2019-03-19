package fantasy

type Projection struct {
	NextRank          int
	NextPoints        float64
	NextDegradeBy     float64
	CalculateNextFunc func()
}

func newProjection(basePoints float64, degradeBy float64) *Projection {

	p := &Projection{
		NextRank:      1,
		NextPoints:    basePoints,
		NextDegradeBy: 0,
	}

	p.CalculateNextFunc = func() {
		p.NextRank++
		p.NextDegradeBy -= degradeBy
		p.NextPoints -= p.NextDegradeBy
	}

	return p
}

func (p *Projection) CalculateNext() {
	p.CalculateNextFunc()
}
