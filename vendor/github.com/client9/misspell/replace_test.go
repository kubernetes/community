package misspell

import (
	"strings"
	"testing"
)

func TestReplaceIgnore(t *testing.T) {
	cases := []struct {
		ignore string
		text   string
	}{
		{"knwo,gae", "https://github.com/Unknwon, github.com/hnakamur/gaesessions"},
	}
	for line, tt := range cases {
		r := New()
		r.RemoveRule(strings.Split(tt.ignore, ","))
		r.Compile()
		got, _ := r.Replace(tt.text)
		if got != tt.text {
			t.Errorf("%d: Replace files want %q got %q", line, tt.text, got)
		}
	}
}

func TestReplaceLocale(t *testing.T) {
	cases := []struct {
		orig string
		want string
	}{
		{"The colours are pretty", "The colors are pretty"},
		{"summaries", "summaries"},
	}

	r := New()
	r.AddRuleList(DictAmerican)
	r.Compile()
	for line, tt := range cases {
		got, _ := r.Replace(tt.orig)
		if got != tt.want {
			t.Errorf("%d: ReplaceLocale want %q got %q", line, tt.orig, got)
		}
	}
}

func TestReplace(t *testing.T) {
	cases := []struct {
		orig string
		want string
	}{
		{"I live in Amercia", "I live in America"},
		{"grill brocoli now", "grill broccoli now"},
		{"There is a zeebra", "There is a zebra"},
		{"foo other bar", "foo other bar"},
		{"ten fiels", "ten fields"},
		{"Closeing Time", "Closing Time"},
		{"closeing Time", "closing Time"},
		{" TOOD: foobar", " TODO: foobar"},
		{" preceed ", " precede "},
		{"preceeding", "preceding"},
		{"functionallity", "functionality"},
	}
	r := New()
	for line, tt := range cases {
		got, _ := r.Replace(tt.orig)
		if got != tt.want {
			t.Errorf("%d: Replace files want %q got %q", line, tt.orig, got)
		}
	}
}

func TestCheckReplace(t *testing.T) {
	r := Replacer{
		engine: NewStringReplacer("foo", "foobar", "runing", "running"),
		corrected: map[string]string{
			"foo":    "foobar",
			"runing": "running",
		},
	}

	s := "nothing at all"
	news, diffs := r.Replace(s)
	if s != news || len(diffs) != 0 {
		t.Errorf("Basic recheck failed: %q vs %q", s, news)
	}

	//
	// Test single, correct,.Correctedacements
	//
	s = "foo"
	news, diffs = r.Replace(s)
	if news != "foobar" || len(diffs) != 1 || diffs[0].Original != "foo" && diffs[0].Corrected != "foobar" && diffs[0].Column != 0 {
		t.Errorf("basic recheck1 failed %q vs %q", s, news)
	}
	s = "foo junk"
	news, diffs = r.Replace(s)
	if news != "foobar junk" || len(diffs) != 1 || diffs[0].Original != "foo" && diffs[0].Corrected != "foobar" && diffs[0].Column != 0 {
		t.Errorf("basic recheck2 failed %q vs %q, %v", s, news, diffs[0])
	}

	s = "junk foo"
	news, diffs = r.Replace(s)
	if news != "junk foobar" || len(diffs) != 1 || diffs[0].Original != "foo" && diffs[0].Corrected != "foobar" && diffs[0].Column != 5 {
		t.Errorf("basic recheck3 failed: %q vs %q", s, news)
	}

	s = "junk foo junk"
	news, diffs = r.Replace(s)
	if news != "junk foobar junk" || len(diffs) != 1 || diffs[0].Original != "foo" && diffs[0].Corrected != "foobar" && diffs[0].Column != 5 {
		t.Errorf("basic recheck4 failed: %q vs %q", s, news)
	}

	// Incorrect.Correctedacements
	s = "food pruning"
	news, _ = r.Replace(s)
	if news != s {
		t.Errorf("incorrect.Correctedacement failed: %q vs %q", s, news)
	}
}
