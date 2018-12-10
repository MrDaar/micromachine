package examples_test

import (
	"testing"

	mm "github.com/MrDaar/micromachine"
	e "github.com/MrDaar/micromachine/examples"
)

func TestNewVehicle(t *testing.T) {
	v := e.NewVehicle(
		e.WithState(e.StateReady),
		e.WithStatePaths(defaultStatePaths()),
	)
	err := v.SetState(mm.GroupPublic, e.StateRiding)
	if err != nil {
		t.Fatal(err.Error())
	}

	if got := v.GetState(); got != e.StateRiding {
		t.Fatalf(`state error: got %d, want %d`, got, e.StateRiding)
	}
}

func defaultStatePaths() mm.Paths {
	return mm.Paths{
		mm.GroupPublic: {
			e.StateReady: {e.StateRiding: {
				func() error {
					return nil
				},
			}},
		},
	}
}
