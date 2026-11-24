package main

import (
	"reflect"
	"testing"
)

func TestZigzag(t *testing.T) {
	root := tn(3, tn(9, nil, nil), tn(20, tn(15, nil, nil), tn(7, nil, nil)))
	want := [][]int{{3}, {20, 9}, {15, 7}}
	if got := Zigzag(root); !reflect.DeepEqual(got, want) {
		t.Errorf("Zigzag = %v, want %v", got, want)
	}

	full := tn(1, tn(2, tn(4, nil, nil), tn(5, nil, nil)), tn(3, tn(6, nil, nil), tn(7, nil, nil)))
	want2 := [][]int{{1}, {3, 2}, {4, 5, 6, 7}}
	if got := Zigzag(full); !reflect.DeepEqual(got, want2) {
		t.Errorf("Zigzag = %v, want %v", got, want2)
	}

	if got := Zigzag(nil); !reflect.DeepEqual(got, [][]int{}) {
		t.Errorf("Zigzag(nil) = %v", got)
	}
}
