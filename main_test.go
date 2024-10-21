package main

import (
	"reflect"
	"testSorts/sorts"
	"testing"
)

var tests = []struct {
	input    []int
	expected []int
}{
	{[]int{8, 7, 9, 8, 2, 4, 6, 1}, []int{1, 2, 4, 6, 7, 8, 8, 9}},
	{[]int{0, 0, 0}, []int{0, 0, 0}},
	{[]int{5, 6, 7, 8, 9, 10}, []int{5, 6, 7, 8, 9, 10}},
	{[]int{10, 9, 8, 7, 6, 5}, []int{5, 6, 7, 8, 9, 10}},
	{[]int{}, []int{}},
	{[]int{-5, -9, -8, 0, 4, 1}, []int{-9, -8, -5, 0, 1, 4}},
}

func TestBubbleSort(t *testing.T) {
	for _, test := range tests {
		result := sorts.BubbleSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("[sorts.BubbleSort] For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	for _, test := range tests {
		result := sorts.InsertionSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("[sorts.InsertionSort] For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestMergeSort(t *testing.T) {
	for _, test := range tests {
		result := sorts.MergeSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("[sorts.MergeSort] For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestQuickSort(t *testing.T) {
	for _, test := range tests {
		result := sorts.QuickSort(test.input, 0, len(test.input)-1)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("[sorts.QuickSort] For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	for _, test := range tests {
		result := sorts.SelectionSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("[sorts.SelectionSort] For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}
