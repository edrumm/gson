package gson

import "testing"

func TestParser_Parse(t *testing.T) {

	_, err := Parse(`"language": "Go", "best": true`)

	if err != nil {
		err.PrintError()
		t.Fail()
	}
}

func TestParser_ToString(t *testing.T) {

	obj, err := Parse(`"language": "Go", "best": true`)

	if err != nil {
		err.PrintError()
		t.Fail()
	}

	str := ToString(obj)

	// TODO
	if str != "" {
		t.Fatalf("expected _ but was %s", str)
	}
}
