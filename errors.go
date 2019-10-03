package micromachine

import "errors"

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
