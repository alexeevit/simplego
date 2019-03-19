package api

import (
	"testing"
)

var urlToPatternTests = []struct {
	v    string
	want string
}{
	{"/pets", "/pets"},
	{"/pets/dog", "/pets/{petId}"}}

func TestURLToPattern(t *testing.T) {
	ut := NewURLTransformer()
	ut.LoadPaths([]string{"/pets", "/pets/{petId}"})

	for _, tt := range urlToPatternTests {
		got := ut.URLActualToPattern(tt.v)
		if got != tt.want {
			t.Errorf(`URLActualToPattern("%v") Failed: want [%v], got [%v]`, tt.v, tt.want, got)
		}
	}
}
