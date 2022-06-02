package libganache

import (
	"fmt"
	"testing"
)

func TestGanache(t *testing.T) {
	g, err := RunGanache(nil)
	if err != nil {
		fmt.Println("error running ganache", err)
		t.Fail()
	}

	accs, err := g.Accounts()
	if err != nil {
		fmt.Println("error getting accounts", err)
		t.Fail()
	}

	if len(accs) < 1 {
		fmt.Println("unexpected accounts len")
		t.Fail()
	}

	if g == nil {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}

}
