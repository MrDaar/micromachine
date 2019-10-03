package micromachine_test

import (
	"errors"
	"testing"

	mm "github.com/MrDaar/micromachine"
	mock "github.com/MrDaar/micromachine/test/mock"
)

const (
	// Default.
	StateUnknown mm.State = 0

	// Operational statuses.
	StateReady      mm.State = 10
	StateRiding     mm.State = 11
	StateBatteryLow mm.State = 12
	StateBounty     mm.State = 20
	StateCollected  mm.State = 21
	StateDropped    mm.State = 22

	// Not commissioned for service.
	StateServiceMode mm.State = 30
	StateTerminated  mm.State = 32

	GroupPublic mm.Group = mm.GroupPublic
	GroupUser   mm.Group = 1
	GroupStaff  mm.Group = 2
	GroupAdmin  mm.Group = mm.GroupSuperAdmin
	GroupSystem mm.Group = 4
)

var errTransitionDisabled = errors.New(`transition disabled`)

func TestTransition(t *testing.T) {
	sm := &mm.StateMachine{
		Paths: mm.Paths{
			GroupPublic: {
				StateReady:  {StateRiding: nil},
				StateRiding: {StateReady: nil},
			},
			GroupStaff: {
				StateBounty:    {StateCollected: nil},
				StateCollected: {StateDropped: nil},
				StateDropped:   {StateReady: nil},
			},
			GroupAdmin: {
				StateCollected: {StateDropped: {
					func() error {
						t.Log(errTransitionDisabled.Error())
						return errTransitionDisabled
					},
				}},
			},
			GroupSystem: {
				StateReady:  {StateUnknown: nil, StateBounty: nil},
				StateRiding: {StateBatteryLow: nil},
				StateBatteryLow: {StateBounty: {
					func() error {
						t.Log(`Sending bounty notification...`)
						return nil
					},
				}},
			},
		},
	}

	sm.Paths[GroupSystem][StateRiding][StateBatteryLow] = []mm.Condition{
		func() error {
			t.Log(`Transitioning to bounty state...`)
			sm.Subject = &mock.Stateful{
				GetStateFunc: func() mm.State { return StateBatteryLow },
			}
			return sm.Transition(GroupSystem, StateBounty)
		},
	}

	for i, tt := range []struct {
		in   mm.Group
		from mm.State
		to   mm.State
		want error
	}{
		// Pass
		{GroupPublic, StateReady, StateRiding, nil},
		{GroupPublic, StateRiding, StateReady, nil},
		{GroupUser, StateRiding, StateReady, nil},
		{GroupStaff, StateBounty, StateCollected, nil},
		{GroupStaff, StateCollected, StateDropped, nil},
		{GroupStaff, StateDropped, StateReady, nil},
		{GroupAdmin, StateCollected, StateReady, nil},
		{GroupAdmin, StateBatteryLow, StateDropped, nil},
		{GroupAdmin, StateBounty, StateUnknown, nil},
		{GroupAdmin, StateServiceMode, StateTerminated, nil},
		{GroupSystem, StateReady, StateUnknown, nil},
		{GroupSystem, StateReady, StateBounty, nil},
		{GroupSystem, StateBatteryLow, StateBounty, nil},
		{GroupSystem, StateRiding, StateBatteryLow, nil},

		// Fail
		{123, StateCollected, StateReady, mm.ErrInvalidGroup},
		{GroupPublic, StateCollected, StateReady, mm.ErrInvalidFromState},
		{GroupPublic, StateReady, StateCollected, mm.ErrInvalidToState},
		{GroupAdmin, StateCollected, StateDropped, errTransitionDisabled},
	} {
		in := tt.in
		from := tt.from
		to := tt.to
		want := tt.want

		sm.Subject = &mock.Stateful{
			GetStateFunc: func() mm.State { return from },
		}

		if err := sm.Transition(in, to); err != want {
			if err != nil {
				t.Log(err.Error())
			}
			t.Fatalf(`Test %d - Could not transition - in %d, from %d, to %d`, i, in, from, to)
		}
	}
}
