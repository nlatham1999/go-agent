package sortedset

import (
	"errors"
	"sort"
)

type setKey struct {
	previous interface{}
	current  interface{}
	next     interface{}
}

type SortedSet struct {
	head    interface{}
	tail    interface{}
	values  map[interface{}]*setKey
	pointer *setKey
}

type sortedSetSorter struct {
	elements []interface{}
	f        func(e interface{}) interface{}
	reverse  bool
}

func (s *sortedSetSorter) Len() int {
	return len(s.elements)
}

func (s *sortedSetSorter) Less(i, j int) bool {

	var element1 interface{}
	var element2 interface{}
	if s.reverse {
		element1 = s.f(s.elements[i])
		element2 = s.f(s.elements[j])
	} else {
		element1 = s.f(s.elements[j])
		element2 = s.f(s.elements[i])
	}

	switch element1.(type) {
	case int:
		return element1.(int) > element2.(int)
	case string:
		return element1.(string) > element2.(string)
	case float64:
		return element1.(float64) > element2.(float64)
	default:
		return false
	}
}

func (s *sortedSetSorter) Swap(i int, j int) {
	s.elements[i], s.elements[j] = s.elements[j], s.elements[i]
}

var (
	ErrItemExists      = errors.New("item exists")
	ErrItemDoesntExist = errors.New("item doesn't exist")
)

// create a new sorted set with the given values
func NewSortedSet(values ...interface{}) *SortedSet {
	ss := &SortedSet{
		head:   nil,
		tail:   nil,
		values: make(map[interface{}]*setKey),
	}
	for _, value := range values {
		ss.Add(value)
	}
	if len(values) > 0 {
		ss.pointer = ss.values[ss.head]
	}
	return ss
}

// adds a value to the end of the set
func (ss *SortedSet) Add(value interface{}) error {
	if _, ok := ss.values[value]; ok {
		return ErrItemExists
	}

	key := &setKey{}

	if ss.head == nil { // if the set is empty
		ss.head = value
		ss.tail = value
		key.previous = nil
		key.next = nil
		ss.pointer = key // set the pointer to the new value
	} else {
		key.previous = ss.tail
		key.next = nil
		ss.values[ss.tail].next = value
		ss.tail = value
	}
	key.current = value
	ss.values[value] = key

	return nil
}

// returns the next value in the set after the given value
func (ss *SortedSet) After(value interface{}) (interface{}, error) {
	if key, ok := ss.values[value]; ok {
		return key.next, nil
	}
	return nil, ErrItemDoesntExist
}

// takes a function that returns a boolean and returns true if all values in the set return true
func (ss *SortedSet) All(operation func(value interface{}) bool) bool {
	for value := ss.First(); value != nil; value, _ = ss.Next() {
		loopPointer := ss.pointer
		if !operation(value) {
			return false
		}
		ss.pointer = loopPointer
	}
	return true
}

// takes a function that returns a boolean and returns true if any value in the set returns true
func (ss *SortedSet) Any(operation func(value interface{}) bool) bool {
	for value := ss.First(); value != nil; value, _ = ss.Next() {
		loopPointer := ss.pointer
		if operation(value) {
			return true
		}
		ss.pointer = loopPointer
	}
	return false
}

// takes a function and applies it to all values in the set
func (ss *SortedSet) Ask(operation func(value interface{})) {
	for value := ss.First(); value != nil; value, _ = ss.Next() {
		loopPointer := ss.pointer
		operation(value)
		ss.pointer = loopPointer
	}
}

// returns the value before the given value
func (ss *SortedSet) Before(value interface{}) (interface{}, error) {
	if key, ok := ss.values[value]; ok {
		return key.previous, nil
	}
	return nil, ErrItemDoesntExist
}

// returns if the value exists in the set
func (ss *SortedSet) Contains(value interface{}) bool {
	_, ok := ss.values[value]
	return ok
}

// returns the current value of the pointer
func (ss *SortedSet) Current() interface{} {
	if ss.pointer != nil {
		return ss.pointer.current
	} else {
		return nil
	}
}

// returns the difference of this set and the given set as a new set (this - set)
func (ss *SortedSet) Difference(set *SortedSet) *SortedSet {
	newSet := NewSortedSet()
	for value := ss.First(); value != nil; value, _ = ss.Next() {
		loopPointer := ss.pointer
		if !set.Contains(value) {
			newSet.Add(value)
		}
		ss.pointer = loopPointer
	}
	return newSet
}

// returns if the set is empty
func (ss *SortedSet) Empty() bool {
	return len(ss.values) == 0
}

// returns the first value in the set
func (ss *SortedSet) First() interface{} {
	ss.pointer = ss.values[ss.head]
	return ss.head
}

// inserts a value after another value
func (ss *SortedSet) InsertAfter(value, after interface{}) error {
	if _, ok := ss.values[value]; ok {
		return ErrItemExists
	}

	if after == ss.tail {
		ss.tail = value
	}

	key := &setKey{}

	if afterKey, ok := ss.values[after]; ok {
		key.previous = after
		key.current = value
		key.next = afterKey.next
		afterKey.next = value
		ss.values[value] = key
	} else {
		return ErrItemDoesntExist
	}

	return nil
}

// inserts a value before another value
func (ss *SortedSet) InsertBefore(value, before interface{}) error {
	if _, ok := ss.values[value]; ok {
		return ErrItemExists
	}

	key := &setKey{}

	if _, ok := ss.values[before]; !ok {
		return ErrItemDoesntExist
	}

	if before == ss.head {
		key.previous = nil
		key.next = before
		ss.values[before].previous = value
		ss.head = value
		ss.values[value] = key
		return nil
	}

	ss.values[ss.values[before].previous].next = value
	key.previous = ss.values[before].previous
	ss.values[before].previous = value
	key.next = before
	key.current = value
	ss.values[value] = key

	return nil
}

// returns the intersection of this set and the given set as a new set
func (ss *SortedSet) Intersection(set *SortedSet) *SortedSet {
	newSet := NewSortedSet()
	for value := ss.First(); value != nil; value, _ = ss.Next() {
		loopPointer := ss.pointer
		if set.Contains(value) {
			newSet.Add(value)
		}
		ss.pointer = loopPointer
	}
	return newSet
}

// returns the last value in the set and sets the pointer to it
func (ss *SortedSet) Last() interface{} {
	ss.pointer = ss.values[ss.tail]
	return ss.tail
}

// returns the length of the set
func (ss *SortedSet) Len() int {
	return len(ss.values)
}

// returns the set as a slice
func (ss *SortedSet) List() []interface{} {
	slice := make([]interface{}, 0, len(ss.values))
	for key := ss.head; key != nil; key = ss.values[key].next {
		slice = append(slice, key)
	}
	return slice
}

// moves the pointer to the next value in the set and returns it
func (ss *SortedSet) Next() (interface{}, error) {
	if ss.pointer == nil {
		return nil, ErrItemDoesntExist
	}

	value := ss.pointer.next
	if value == nil {
		return nil, ErrItemDoesntExist
	}

	ss.pointer = ss.values[value]
	return value, nil
}

// moves the pointer to the previous value in the set and returns it
func (ss *SortedSet) Previous() (interface{}, error) {
	if ss.pointer == nil {
		return nil, ErrItemDoesntExist
	}

	value := ss.pointer.previous
	if value == nil {
		return nil, ErrItemDoesntExist
	}

	ss.pointer = ss.values[value]
	return value, nil
}

// returns a random value from the set
func (ss *SortedSet) RandomValue() interface{} {
	for key, _ := range ss.values {
		return key
	}
	return nil
}

// returns a random value from the set that satisfies the given operation
func (ss *SortedSet) RandomValueWhere(operation func(value interface{}) bool) interface{} {
	for key, _ := range ss.values {
		if operation(key) {
			return key
		}
	}
	return nil
}

// removes a value from the set
func (ss *SortedSet) Remove(value interface{}) error {
	if _, ok := ss.values[value]; ok {
		key := ss.values[value]

		if key.previous != nil {
			ss.values[key.previous].next = key.next
		} else {
			ss.head = key.next
		}

		if key.next != nil {
			ss.values[key.next].previous = key.previous
		} else {
			ss.tail = key.previous
		}

		// if the pointer is the value being removed, set the pointer.next
		// key is being removed from the map so it is ok to change it
		if ss.pointer == key {
			ss.pointer.next = key.next
		}

		delete(ss.values, value)
	} else {
		return ErrItemDoesntExist
	}

	if len(ss.values) == 0 {
		ss.head = nil
		ss.tail = nil
	}

	return nil
}

func (ss *SortedSet) sort(reverse bool, f func(e interface{}) interface{}) {
	// convert to list
	list := ss.List()

	// sort the list using the sorting interface
	sorter := &sortedSetSorter{
		elements: list,
		f:        f,
		reverse:  reverse,
	}
	sort.Sort(sorter)

	//reset the sorted set as the list
	ss.head = nil
	ss.tail = nil
	ss.values = map[interface{}]*setKey{}
	for _, value := range list {
		ss.Add(value)
	}

}

// sorts the set in ascending order
func (ss *SortedSet) SortAsc(f func(e interface{}) interface{}) {
	ss.sort(false, f)
}

// sorts the set in descending order
func (ss *SortedSet) SortDesc(f func(e interface{}) interface{}) {
	ss.sort(true, f)
}

// returns the symmetric difference of this set and the given set as a new set (this XOR set)
func (ss *SortedSet) SymmetricDifference(set *SortedSet) *SortedSet {
	newSet := NewSortedSet()
	for value := ss.First(); value != nil; value, _ = ss.Next() {
		loopPointer := ss.pointer
		if !set.Contains(value) {
			newSet.Add(value)
		}
		ss.pointer = loopPointer
	}

	for value := set.First(); value != nil; value, _ = set.Next() {
		loopPointer := set.pointer
		if !ss.Contains(value) {
			newSet.Add(value)
		}
		set.pointer = loopPointer
	}

	return newSet
}

// returns the union of this set and the given set as a new set
func (ss *SortedSet) Union(set *SortedSet) *SortedSet {
	newSet := NewSortedSet()
	for value := ss.First(); value != nil; value, _ = ss.Next() {
		loopPointer := ss.pointer
		newSet.Add(value)
		ss.pointer = loopPointer
	}

	for value := set.First(); value != nil; value, _ = set.Next() {
		loopPointer := set.pointer
		newSet.Add(value)
		set.pointer = loopPointer
	}

	return newSet
}
