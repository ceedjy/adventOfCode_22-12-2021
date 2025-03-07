// Author : Cassiopée Gossin

package main

import (
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	for range b.N {
		pb1(filename)
	}
}

func BenchmarkPart2(b *testing.B) {
	for range b.N {
		pb2(filename)
	}
}
