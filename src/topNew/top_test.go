package top

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTop_Seq(t *testing.T) {
	fmt.Println("Seq TopNew Test...")
	topData, err := New(10).Seq().Top(5)
	if err != nil || !reflect.DeepEqual(topData, []int{10, 9, 8, 7, 6}) {
		t.Errorf("Got %v, expected [10 9 8 7 6]. error: %v", topData, err)
	}

	test := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	topData, err = (*TopData)(&test).Top(5)
	if err != nil || !reflect.DeepEqual(topData, []int{10, 9, 8, 7, 6}) {
		t.Errorf("Got %v, expected [10 9 8 7 6]. error: %v", topData, err)
	}

	test = []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	topData, err = (*TopData)(&test).Top(5)
	if err != nil || !reflect.DeepEqual(topData, []int{2, 2, 2, 2, 2}) {
		t.Errorf("Got %v, expected [2 2 2 2 2]. error: %v", topData, err)
	}
}

func BenchmarkTop(b *testing.B) {
	fmt.Println("TopNew Benchmark Test...")
	p := New(100000).Rand(100000)

	for i := 0; i < b.N; i++ {
		_, err := p.Top(100)
		if err != nil {
			b.Errorf("error: %v", err)
		}
	}
}

func BenchmarkParallelTop(b *testing.B) {
	fmt.Println("TopNew Parallel Benchmark Test...")
	p := New(100000).Rand(100000)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := p.Top(100)
			if err != nil {
				b.Errorf("error: %v", err)
			}
		}
	})
}
