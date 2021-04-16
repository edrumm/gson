package gson

import "testing"

func TestParser_Parse(t *testing.T) {
	parser := NewParser()

	_, err := parser.Parse(`"language": "Go", "best": true`)

	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
}

func TestParser_ToString(t *testing.T) {
	parser := NewParser()

	obj, err := parser.Parse(`"language": "Go", "best": true`)

	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	str := parser.ToString(obj)

	// TODO
	if str != "" {
		t.Fatalf("expected _ but was %s", str)
	}
}
