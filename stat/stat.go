package stat

// Statter provides the interface for measuring quantifiable events
type Statter interface {
	Count(stat string, i int) error
	Measure(stat string, v float64) error
}
