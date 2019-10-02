package micromachine

import "errors"

const (
	// GroupPublic allows anyone to make the transitions in this group.
	GroupPublic Group = 0

	// GroupSuperAdmin bypasses all map checks, but still runs conditions.
	GroupSuperAdmin Group = 99
)

var (
	// ErrSubjectCannotBeNil occurs when subject is nil.
	ErrSubjectCannotBeNil = errors.New(`subject cannot be nil`)

	// ErrInvalidGroup occurs when a group does not exist.
	ErrInvalidGroup = errors.New(`invalid group`)

	// ErrInvalidFromState occurs when a from state does not exist.
	ErrInvalidFromState = errors.New(`"from" state is invalid`)

	// ErrInvalidToState occurs when a to state does not exist.
	ErrInvalidToState = errors.New(`"to" state is invalid`)
)

// State is an int that represents a state.
type State int

// Group is an int that represents a group.
// Groups can have different rules.
type Group int

// Condition is a function that must be checked before transitioning.
// Can also be used as a callback.
type Condition func() error

// Paths maps the possible routes in the machine.
type Paths map[Group]map[State]map[State][]Condition

// StateMachine is a wrapper for the paths,
// and has a reference to that which we are managing state for.
type StateMachine struct {
	Paths   Paths
	Subject Stateful
}

// Stateful gives us a way to interact with our subject.
type Stateful interface {
	GetState() State
	SetState(Group, State) error
}

// Transition validates state transitions.
func (sm *StateMachine) Transition(in Group, to State) error {
	if sm.Subject == nil {
		return ErrSubjectCannotBeNil
	}

	if in != GroupPublic {
		if err := sm.Transition(GroupPublic, to); err == nil {
			return nil
		}
	}

	pathIn, ok := sm.Paths[in]
	if !ok && in != GroupSuperAdmin {
		return ErrInvalidGroup
	}

	pathFrom, ok := pathIn[sm.Subject.GetState()]
	if !ok && in != GroupSuperAdmin {
		return ErrInvalidFromState
	}

	pathTo, ok := pathFrom[to]
	if !ok && in != GroupSuperAdmin {
		return ErrInvalidToState
	}

	// could be done in channels like: https://github.com/ryanfaerman/fsm/blob/master/fsm.go#L59
	// but you may want the checks to depend on one another
	for _, c := range pathTo {
		if err := c(); err != nil {
			return err
		}
	}

	return nil
}
