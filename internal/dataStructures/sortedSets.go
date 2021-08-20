package dataStructures

import (
	"sort"
)

type SortedSet struct {
	data []int
}

func NewSortedSet(elements []int) *SortedSet {
	l := SortedSet{data: make([]int, 0)}
	for _, e := range elements {
		if !l.contains(e) {
			l.data = append(l.data, e)
		}
	}
	sort.Ints(l.data)
	return &l
}

// Add adds an element to the set. If it already exists, nothing
// happens, as set elements are unique.
func (l *SortedSet) Add(element int) {
	if len(l.data) == 0 {
		l.data = []int{element}
		return
	}

	if l.contains(element) {
		return
	}

	l.data = append(l.data, element)
	sort.Ints(l.data)
}

// DeleteMax removes the last element
func (l *SortedSet) DeleteMax() int {

	if len(l.data) == 1 {
		elem := l.data[0]
		l.data = make([]int, 0)
		return elem
	}

	elem := l.data[len(l.data)-1]
	l.data = l.data[:len(l.data)-2]
	return elem
}

// DeleteMin removes the first element
func (l *SortedSet) DeleteMin() int {

	if len(l.data) == 1 {
		elem := l.data[0]
		l.data = make([]int, 0)
		return elem
	}
	elem := l.data[0]
	l.data = l.data[1:]
	return elem
}

// Contains reports if the set contains the given element.
func (l *SortedSet) contains(search int) bool {
	for _, e := range l.data {
		if search == e {
			return true
		}
	}
	return false
}

func (l *SortedSet) GetMin() int {
	return l.data[0]
}

func (l *SortedSet) GetMax() int {
	return l.data[len(l.data)-1]
}
