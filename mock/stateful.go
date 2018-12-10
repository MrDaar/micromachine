package mock

import mm "github.com/MrDaar/micromachine"

// Stateful is used to mock mm.Stateful
type Stateful struct {
	GetStateFunc func() mm.State
	SetStateFunc func(mm.Group, mm.State) error
}

// GetState is used to mock mm.Stateful.GetState
func (s *Stateful) GetState() mm.State {
	if s.GetStateFunc == nil {
		panic("*Stateful.GetState was called, but it is not mocked")
	}

	return s.GetStateFunc()
}

// SetState is used to mock mm.Stateful.SetState
func (s *Stateful) SetState(in mm.Group, to mm.State) error {
	if s.SetStateFunc == nil {
		panic("*Stateful.SetState was called, but it is not mocked")
	}

	return s.SetStateFunc(in, to)
}
