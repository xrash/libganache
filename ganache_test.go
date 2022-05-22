package libganache

import (
	"testing"
)

func TestGanache(t *testing.T) {
	g, err := RunGanache(nil)

	if g == nil {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}
}
