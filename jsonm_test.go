package jsonm

import "testing"

var jsonData = []byte(`{"one":{"two":123,"two1":"string","two2":{"three":3.3,"three1":[3,6,7,8]}}}`)

func TestNew(t *testing.T) {
	jm, err := NewJsonM(jsonData)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(jm)
}

func TestMap(t *testing.T) {
	jm, err := NewJsonM(jsonData)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(jm.Map("one"))

	t.Log(jm.JsonM("one"))

	t.Log(jm.Maps())

	t.Log(jm.Maps("one", "two"))

	t.Log(jm.Maps("one", "two2", "three1"))
}

func TestJsonMaps(t *testing.T) {
	jm, ok, err := JsonMap(jsonData, "one")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(jm, ok)

	jm, ok, err = JsonMaps(jsonData, "one", "two")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(jm, ok)

	jm, ok, err = JsonMaps(jsonData, "one", "two2", "three")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(jm, ok)
}
