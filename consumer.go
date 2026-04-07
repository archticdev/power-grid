package power_grid

import "time"

type Consumer interface {
	// Consume returns the amount of energy consumed in kWh for the given duration and time of day.
	Consume(duration time.Duration, timeOfDay time.Time) float64
}

type simpleConsumer struct {
	consumption float64
}

// NewSimpleConsumer creates a new Consumer with a fixed consumption rate in kW.
func NewSimpleConsumer(consumption float64) Consumer {
	return simpleConsumer{consumption: consumption}
}

func (p simpleConsumer) Consume(duration time.Duration, _ time.Time) float64 {
	return p.consumption * duration.Hours()
}
