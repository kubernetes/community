package misspell

import (
	"testing"
)

func TestNotWords(t *testing.T) {
	cases := []struct {
		word string
		want string
	}{
		{" /foo/bar abc", "          abc"},
		{"X/foo/bar abc", "X/foo/bar abc"},
		{"[/foo/bar] abc", "[        ] abc"},
		{"/", "/"},
		{"x nickg@client9.xxx y", "x                   y"},
		{"x infinitie.net y", "x               y"},
		{"(s.svc.GetObject(", "(               ("},
		{"\\nto", "  to"},
	}
	for pos, tt := range cases {
		got := RemoveNotWords(tt.word)
		if got != tt.want {
			t.Errorf("%d want %q  got %q", pos, tt.want, got)
		}
	}
}
