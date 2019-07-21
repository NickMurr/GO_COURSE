package main

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}
func BenchmarkBigLen(b *testing.B) {
	big := NewBig()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Len()
	}
}
