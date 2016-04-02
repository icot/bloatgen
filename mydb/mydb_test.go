package mydb

import (
    "testing"
    "reflect"
)

func TestGen(t *testing.T) {
	for c := 0; c < 10; c++ {
		got := GenerateData()
		if len(got) <= 32 {
			t.Errorf("GenerateData() lenght <= 32, want %q", got)
		}
        if reflect.TypeOf(got) != reflect.TypeOf("string") {
			t.Errorf("GenerateData() is not a string")
        }
	}
}
