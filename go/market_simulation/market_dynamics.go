package market_simulation

import (
	"github.com/umbralcalc/stochadex/pkg/phenomena"
	"github.com/umbralcalc/stochadex/pkg/simulator"
)

// MarketDynamicsIteration defines an iteration for a dynamical model of financial market
// microstructure which was defined in this chapter of the book Worlds Of Observation:
// https://umbralcalc.github.io/worlds-of-observation/algo_trading_on_financial_markets/chapter.pdf
type MarketDynamicsIteration struct {
	coxProcess *phenomena.CoxProcessIteration
}

func (m *MarketDynamicsIteration) Configure(
	partitionIndex int,
	settings *simulator.LoadSettingsConfig,
) {
	m.coxProcess.Configure(partitionIndex, settings)
}

func (m *MarketDynamicsIteration) Iterate(
	otherParams *simulator.OtherParams,
	partitionIndex int,
	stateHistories []*simulator.StateHistory,
	timestepsHistory *simulator.CumulativeTimestepsHistory,
) []float64 {
	// do something
	return []float64{}
}
