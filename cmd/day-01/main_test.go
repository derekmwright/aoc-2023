package main

import (
	"os"
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for i := 0; i < b.N; i += 1 {
		_, _ = Part1(f)
	}

	b.ReportAllocs()
}
