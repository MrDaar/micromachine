package mock

import (
	"fmt"

	mm "github.com/MrDaar/micromachine"
)

const undefinedStatefulFuncMessage = `called %[1]s but %[1]sFunc is not defined in Stateful`

var _ mm.Stateful = (*Stateful)(nil)

// Stateful implements mm.Stateful.
type Stateful struct {
	GetStateFunc func() mm.State
	SetStateFunc func(mm.Group, mm.State) error
}

func (s *Stateful) GetState() mm.State {
	if s.GetStateFunc == nil {
		panic(fmt.Sprintf(undefinedStatefulFuncMessage, "GetState"))
	}
	return s.GetStateFunc()
}

func (s *Stateful) SetState(in mm.Group, to mm.State) error {
	if s.SetStateFunc == nil {
		panic(fmt.Sprintf(undefinedStatefulFuncMessage, "SetState"))
	}
	return s.SetStateFunc(in, to)
}
