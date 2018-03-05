package misspell

import (
	"bytes"
	"io/ioutil"
	"testing"
)

var (
	sampleClean string
	sampleDirty string
	tmpCount    int
	tmp         string
	rep         *Replacer
)

func init() {

	buf := bytes.Buffer{}
	for i := 0; i < len(DictMain); i += 2 {
		buf.WriteString(DictMain[i+1] + " ")
		if i%5 == 0 {
			buf.WriteString("\n")
		}
	}
	sampleClean = buf.String()
	sampleDirty = sampleClean + DictMain[0] + "\n"
	rep = New()
}

// BenchmarkCleanString takes a clean string (one with no errors)
func BenchmarkCleanString(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	var updated string
	var diffs []Diff
	var count int
	for n := 0; n < b.N; n++ {
		updated, diffs = rep.Replace(sampleClean)
		count += len(diffs)
	}

	// prevent compilier optimizations
	tmpCount = count
	tmp = updated
}

func discardDiff(_ Diff) {
	tmpCount++
}

// BenchmarkCleanStream takes a clean reader (no misspells) and outputs to a buffer
func BenchmarkCleanStream(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	tmpCount = 0
	buf := bytes.NewBufferString(sampleClean)
	out := bytes.NewBuffer(make([]byte, 0, len(sampleClean)+100))
	for n := 0; n < b.N; n++ {
		buf.Reset()
		buf.WriteString(sampleClean)
		out.Reset()
		rep.ReplaceReader(buf, out, discardDiff)
	}
}

// BenchmarkCleanStreamDiscard takes a clean reader and discards output
func BenchmarkCleanStreamDiscard(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	buf := bytes.NewBufferString(sampleClean)
	tmpCount = 0
	for n := 0; n < b.N; n++ {
		buf.Reset()
		buf.WriteString(sampleClean)
		rep.ReplaceReader(buf, ioutil.Discard, discardDiff)
	}
}

// BenchmarkCleanString takes a clean string (one with no errors)
func BenchmarkDirtyString(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	var updated string
	var diffs []Diff
	var count int
	for n := 0; n < b.N; n++ {
		updated, diffs = rep.Replace(sampleDirty)
		count += len(diffs)
	}

	// prevent compilier optimizations
	tmpCount = count
	tmp = updated
}

func BenchmarkCompile(b *testing.B) {
	r := New()
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r.Compile()
	}
}
