package examples

import mm "github.com/MrDaar/micromachine"

const (
	// StateReady it's ready.
	StateReady mm.State = 1

	// StateRiding it's riding.
	StateRiding mm.State = 2
)

var _ mm.Stateful = (*Vehicle)(nil)

// Vehicle is a little example.
type Vehicle struct {
	state mm.State
	sm    *mm.StateMachine
}

// GetState satisfies mm.GetState
func (s *Vehicle) GetState() mm.State {
	return s.state
}

// SetState satisfies mm.SetState
func (s *Vehicle) SetState(in mm.Group, to mm.State) error {
	err := s.sm.Transition(in, to)
	if err != nil {
		return err
	}

	s.state = to
	return nil
}

// NewVehicle creates a new vehicle.
func NewVehicle(options ...func(*Vehicle)) *Vehicle {
	s := &Vehicle{}
	s.sm = &mm.StateMachine{Subject: s}

	for _, f := range options {
		f(s)
	}
	return s
}

// WithState allows setting the initial state.
func WithState(initial mm.State) func(*Vehicle) {
	return func(s *Vehicle) {
		s.state = initial
	}
}

// WithStatePaths allows setting the state rules.
func WithStatePaths(paths mm.Paths) func(*Vehicle) {
	return func(s *Vehicle) {
		s.sm.Paths = paths
	}
}
