package stringutil

import "testing"

func TestReverse(t *testing.T) {
	const in, want = "Hello, world", "dlrow ,olleH"
	got := Reverse(in)
	if got != want {
		t.Errorf("Reverse(%q) == %q, want %q", in, got, want)
	}
}

func TestTableReverse(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	} {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
