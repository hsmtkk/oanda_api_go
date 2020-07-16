package test

import "testing"

func AssertNil(t *testing.T, v interface{}) {
	t.Helper()
	if v != nil {
		t.Errorf("want nil, got %v", v)
	}
}

func AssertEqualString(t *testing.T, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func AssertEqualStrings(t *testing.T, want, got []string) {
	t.Helper()
	if len(want) != len(got) {
		t.Errorf("want %d length, got %d length", len(want), len(got))
	}
	for i := range want {
		if want[i] != got[i] {
			t.Errorf("want %s, got %s", want[i], got[i])
		}
	}
}
