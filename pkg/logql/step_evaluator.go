package logql

import (
	"github.com/prometheus/prometheus/promql"
)

type StepResult interface {
	SampleVector() promql.Vector
	QuantileSketchVec() ProbabilisticQuantileVector
}

type SampleVector promql.Vector

var _ StepResult = SampleVector{}

func (p SampleVector) SampleVector() promql.Vector {
	return promql.Vector(p)
}

func (p SampleVector) QuantileSketchVec() ProbabilisticQuantileVector {
	return ProbabilisticQuantileVector{}
}

// StepEvaluator evaluate a single step of a query.
type StepEvaluator interface {
	// while Next returns a promql.Value, the only acceptable types are Scalar and Vector.
	Next() (ok bool, ts int64, r StepResult)
	// Close all resources used.
	Close() error
	// Reports any error
	Error() error
	// Explain returns a print of the step evaluation tree
	Explain(Node)
}

type EmptyEvaluator struct{}

var _ StepEvaluator = EmptyEvaluator{}

// Close implements StepEvaluator.
func (EmptyEvaluator) Close() error { return nil }

// Error implements StepEvaluator.
func (EmptyEvaluator) Error() error { return nil }

// Next implements StepEvaluator.
func (EmptyEvaluator) Next() (ok bool, ts int64, r StepResult) {
	return false, 0, nil
}
