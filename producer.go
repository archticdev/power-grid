package power_grid

import "time"

type Producer interface {
	// Produce returns the amount of energy produced in kWh for the given duration and time of day.
	Produce(duration time.Duration, timeOfDay time.Time) float64
}

type simpleProducer struct {
	productionCapacity float64
}

// NewSimpleProducer creates a new Producer with a fixed production rate in kW.
func NewSimpleProducer(productionCapacity float64) Producer {
	return simpleProducer{productionCapacity: productionCapacity}
}

func (p simpleProducer) Produce(duration time.Duration, _ time.Time) float64 {
	return p.productionCapacity * duration.Hours()
}
