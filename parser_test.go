package gson

import "testing"

func TestParser_Parse(t *testing.T) {
	parser := NewParser()

	_, err := parser.Parse(`"language": "Go", "best": true`)

	if err != nil {
		err.PrintError()
		t.Fail()
	}
}

func TestParser_ToString(t *testing.T) {
	parser := NewParser()

	obj, err := parser.Parse(`"language": "Go", "best": true`)

	if err != nil {
		err.PrintError()
		t.Fail()
	}

	str := parser.ToString(obj)

	// TODO
	if str != "" {
		t.Fatalf("expected _ but was %s", str)
	}
}
