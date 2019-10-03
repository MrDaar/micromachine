package micromachine

const (
	// GroupPublic allows anyone to make the transitions in this group.
	GroupPublic Group = 0

	// GroupSuperAdmin bypasses all map checks but still runs conditions.
	GroupSuperAdmin Group = 99
)

type (
	// State is an int that represents a state.
	State int

	// Group is an int that represents a group.
	Group int

	// Condition must be checked before transitioning. Can also be used as a callback.
	Condition func() error

	// Paths maps the possible routes in the machine.
	Paths map[Group]map[State]map[State][]Condition
)

// Stateful gives us a way to interact with our subject.
type Stateful interface {
	GetState() State
	SetState(Group, State) error
}

// StateMachine wraps the paths and has a reference to that which we are managing state for.
type StateMachine struct {
	Paths   Paths
	Subject Stateful
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
	for _, f := range pathTo {
		if err := f(); err != nil {
			return err
		}
	}

	return nil
}
