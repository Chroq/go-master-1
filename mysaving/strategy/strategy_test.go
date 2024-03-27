package strategy_test

import (
	"testing"

	"github.com/Chroq/mysaving/strategy"
	"github.com/stretchr/testify/assert"
)

func TestNewStrategy(t *testing.T) {
	t.Run("TestNewStrategy", func(t *testing.T) {
		t.Parallel()
		strategy := strategy.NewStrategy("Classic", strategy.Simulation{})
		assert.Equal(t, strategy.Name(), "Classic")
	})
}

func TestProgressiveStrategy_CalculateSaving(t *testing.T) {
	t.Run("TestProgressiveStrategy_CalculateSaving", func(t *testing.T) {
		t.Parallel()
		simulation := strategy.Simulation{
			Salary: 4000,
			Year:   5,
		}
		strategy := strategy.ProgressiveStrategy{
			Strategy: strategy.NewStrategy("Progressive", simulation),
			Rate:     strategy.ProgressiveRate,
		}
		assert.Equal(t, 48600.0, strategy.CalculateSaving())
	})
}

func TestClassicalStrategy_CalculateSaving(t *testing.T) {
	t.Run("TestClassicalStrategy_CalculateSaving", func(t *testing.T) {
		t.Parallel()
		simulation := strategy.Simulation{
			Salary: 4000,
			Year:   5,
		}
		strategy := strategy.ClassicalStrategy{
			Strategy: strategy.NewStrategy("Classic", simulation),
			Rate:     0.15,
		}
		assert.Equal(t, strategy.CalculateSaving(), 36000.0)
	})
}

func TestPokerStrategy_CalculateSaving(t *testing.T) {
	t.Run("TestPokerStrategy_CalculateSaving", func(t *testing.T) {
		t.Parallel()
		simulation := strategy.Simulation{
			Salary: 4000,
			Year:   5,
		}
		strategy := strategy.PokerStrategy{
			Strategy: strategy.NewStrategy("Poker", simulation),
			MinRate:  strategy.PokerMinRate,
			MaxRate:  strategy.PokerMaxRate,
		}
		assert.Greater(t, strategy.CalculateSaving(), 0.0)
		assert.Less(t, strategy.CalculateSaving(), 1000000.0)
	})
}
