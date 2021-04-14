package gson

import "testing"

func TestParse(t *testing.T) {
	_, err := Parse("")

	if err != nil {
		t.Fail()
	}
}
