package stat

import (
	"github.com/e-dard/gev"
	"github.com/stathat/go"
	"sync"
)

// StatHat implements the Statter interface for the Stat Hat service
type StatHat struct {
	Key string `env:"SH_KEY"`
	mu  sync.RWMutex
}

var std = NewStatHat("")

// NewStatHat returns a new StatHat type.
//
// If key is not the empty string, then it will be used as the Stat Hat
// API key. Otherwise, NewStatHat will attempt to read the Stat Hat API
// key from the environment, looking for a SH_KEY variable.
//
// NewStatHat panics if there is a problem reading this variable, though
// it won't panic if the variable is missing from the environment.
func NewStatHat(key string) StatHat {
	s := StatHat{}
	if err := gev.Unmarshal(&s); err != nil {
		panic(err)
	}

	if key != "" {
		s.Key = key
	}
	return s
}

// Count increments a Stat Hat counter by n. It's threadsafe, and will
// not make a call if the Stat Hat API key is not present.
func (s StatHat) Count(stat string, n int) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.Key == "" {
		return nil
	}
	return stathat.PostEZCount(stat, s.Key, n)
}

// Measure sends a value to a  Stat Hat value. It's threadsafe, and will
// not make a call if the Stat Hat API key is not present.
func (s StatHat) Measure(stat string, v float64) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.Key == "" {
		return nil
	}
	return stathat.PostEZValue(stat, s.Key, v)
}

func Count(stat string, n int) error {
	return std.Count(stat, n)
}

func Measure(stat string, v float64) error {
	return std.Measure(stat, v)
}
