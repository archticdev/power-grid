package power_grid

import "time"

type Transformer interface {
	// Transform returns the net energy resulting from the transformer.
	// An error is returned if the transformer fails, e.g. overload.
	Transform(kWh float64, duration time.Duration) (float64, error)
}

type simpleTransformer struct {
	efficiency float64
}

// NewSimpleTransformer creates a simple Transformer with a fixed percentage of energy successfully transformed.
func NewSimpleTransformer(efficiency float64) Transformer {
	return simpleTransformer{efficiency: efficiency}
}

func (s simpleTransformer) Transform(kWh float64, _ time.Duration) (float64, error) {
	return kWh * s.efficiency, nil
}
