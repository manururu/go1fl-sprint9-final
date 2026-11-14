package main

import (
	"fmt"
	"testing"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-1, 0, 1, 19} {
		t.Run(fmt.Sprint(size), func(t *testing.T) {
			data := generateRandomElements(size)
			want := size
			if want < 0 {
				want = 0
			}
			if len(data) != want {
				t.Fatalf("length = %d, want %d", len(data), want)
			}
			for _, value := range data {
				if value <= 0 {
					t.Fatalf("nonpositive value: %d", value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	cases := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0}, {"empty", []int{}, 0}, {"single", []int{7}, 7},
		{"equal", []int{5, 5, 5}, 5}, {"first", []int{9, 2, 3}, 9},
		{"middle", []int{2, 9, 3}, 9}, {"last", []int{2, 3, 9}, 9},
		{"maxInt", []int{1, int(^uint(0) >> 1), 2}, int(^uint(0) >> 1)},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := maximum(tc.data); got != tc.want {
				t.Fatalf("maximum = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	if got := maxChunks(nil); got != 0 {
		t.Fatalf("nil: got %d", got)
	}
	if got := maxChunks([]int{}); got != 0 {
		t.Fatalf("empty: got %d", got)
	}
	for _, size := range []int{1, 7, 8, 16, 19} {
		for position := 0; position < size; position++ {
			t.Run(fmt.Sprintf("size_%d/max_%d", size, position), func(t *testing.T) {
				data := make([]int, size)
				for i := range data {
					data[i] = 1
				}
				data[position] = 99
				got := maxChunks(data)
				if got != 99 || got != maximum(data) {
					t.Fatalf("maximum = %d, want 99", got)
				}
			})
		}
	}
}
