package sstl

import (
	"unsafe"
)

// Tags: Vector, Array
type Vector[T any] struct {
	// items  []T
	// length int
	data        unsafe.Pointer
	length      int
	capacity    int
	elementSize uintptr
}

func NewVector[T any]() *Vector[T] {
	// Initialise an empty element to get the size
	var zero T
	size := unsafe.Sizeof(zero)
	// Initialise a Vector
	return &Vector[T]{
		data:        nil,
		length:      0,
		capacity:    0,
		elementSize: size,
	}
}

func (s *Vector[T]) Len() int {
	return s.length
}

// func (s *Vector[T]) Insert(item T) bool {

// }
