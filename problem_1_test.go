package main

import "testing"

func TestSolution1(t *testing.T) {
	pairs := [][]int{{10, 23}, {1000, 233168}}
	for _, pair := range pairs {
		in, out := pair[0], pair[1]
		if x := Solution1(in); x != out {
			t.Errorf("Solution1(%d) = %d, should be %d", in, x, out)
		}
	}
}

func TestSolution2(t *testing.T) {
	pairs := [][]int{{10, 23}, {1000, 233168}}
	for _, pair := range pairs {
		in, out := pair[0], pair[1]
		if x := Solution2(in); x != out {
			t.Errorf("Solution2(%d) = %d, should be %d", in, x, out)
		}
	}
}

func BenchmarkSolution1(b *testing.B) {
	in := 1000000000
	for i := 0; i < b.N; i++ {
		Solution1(in)
	}
}

func BenchmarkSolution2(b *testing.B) {
	in := 1000000000
	for i := 0; i < b.N; i++ {
		Solution2(in)
	}
}
