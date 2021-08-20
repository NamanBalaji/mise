package dataStructures

import (
	"sort"
)

type SortedSet struct {
	data []float64
}

func NewSortedSet(elements []float64) *SortedSet {
	l := SortedSet{data: make([]float64, 0)}
	for _, e := range elements {
		if !l.contains(e) {
			l.data = append(l.data, e)
		}
	}
	sort.Float64s(l.data)
	return &l
}

// Add adds an element to the set. If it already exists, nothing
// happens, as set elements are unique.
func (l *SortedSet) Add(element float64) {
	if len(l.data) == 0 {
		l.data = []float64{element}
		return
	}

	if l.contains(element) {
		return
	}

	l.data = append(l.data, element)
	sort.Float64s(l.data)
}

// DeleteMax removes the last element
func (l *SortedSet) DeleteMax() float64 {

	if len(l.data) == 1 {
		elem := l.data[0]
		l.data = make([]float64, 0)
		return elem
	}

	elem := l.data[len(l.data)-1]
	l.data = l.data[:len(l.data)-2]
	return elem
}

// DeleteMin removes the first element
func (l *SortedSet) DeleteMin() float64 {

	if len(l.data) == 1 {
		elem := l.data[0]
		l.data = make([]float64, 0)
		return elem
	}
	elem := l.data[0]
	l.data = l.data[1:]
	return elem
}

// Contains reports if the set contains the given element.
func (l *SortedSet) contains(search float64) bool {
	for _, e := range l.data {
		if search == e {
			return true
		}
	}
	return false
}

func (l *SortedSet) GetMin() float64 {
	return l.data[0]
}

func (l *SortedSet) GetMax() float64 {
	return l.data[len(l.data)-1]
}

func (l *SortedSet) Size() int {
	return len(l.data)
}
