package stringutil

import "testing"

func BenchmarkReverse(b *testing.B) {
	const in = "Hello, world"
	for n := 0; n < b.N; n++ {
		Reverse(in)
	}
}
