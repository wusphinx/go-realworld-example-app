package handy

import (
	"testing"
)

func FuzzUnsafe(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := ToUnsafeBytes(orig)
		doubleRev := ToUnsafeString(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
	})
}

func BenchmarkUnsafe(b *testing.B) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			_ = ToUnsafeBytes(tc)
		}
	}
}

func BenchmarkSafe(b *testing.B) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			_ = []byte(tc)
		}
	}
}
