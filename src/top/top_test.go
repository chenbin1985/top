package top

import (
	"fmt"
	"reflect"
	"testing"
)

// reflect.DeepEqual
func compareSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func TestTop_Case1(t *testing.T) {
	fmt.Println("TestCase1...")
	test := []int{9, 0, 1, 5, 2, 9, 2, 1, 2, 4}
	data, err := (*TopData)(&test).Top(4)
	if err != nil || !compareSlice(data, []int{9, 9, 5, 4}) {
		t.Errorf("Got %v, expected [9 9 5 4]. error: %v", data, err)
	}

}

func TestTop_Seq(t *testing.T) {
	fmt.Println("Seq Top Test...")
	data, err := New(10).Seq().Top(5)
	if err != nil || !reflect.DeepEqual(data, []int{10, 9, 8, 7, 6}) {
		t.Errorf("Got %v, expected [10 9 8 7 6]. error: %v", data, err)
	}

	test := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	data, err = (*TopData)(&test).Top(5)
	if err != nil || !compareSlice(data, []int{10, 9, 8, 7, 6}) {
		t.Errorf("Got %v, expected [10 9 8 7 6]. error: %v", data, err)
	}

	test = []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	data, err = (*TopData)(&test).Top(5)
	if err != nil || !compareSlice(data, []int{2, 2, 2, 2, 2}) {
		t.Errorf("Got %v, expected [2 2 2 2 2]. error: %v", data, err)
	}
}
