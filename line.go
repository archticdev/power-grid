package power_grid

import "time"

type Line interface {
	// Transfer returns the net energy transferred after line loss, given an input over a duration.
	// An error is returned if the line fails to transfer, e.g. a circuit breaker is tripped.
	Transfer(kWh float64, duration time.Duration) (float64, error)
}

type simpleLine struct {
	efficiency float64
}

// NewSimpleLine creates a simple Line with a fixed percentage of energy successfully transferred.
func NewSimpleLine(efficiency float64) Line {
	return simpleLine{efficiency: efficiency}
}

func (s simpleLine) Transfer(kWh float64, _ time.Duration) (float64, error) {
	return kWh * s.efficiency, nil
}
