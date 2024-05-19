package sliceutils

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	sd := &SliceData[int]{}
	index := sd.Push(1, 2, 3)
	expectedData := []int{1, 2, 3}
	expectedSize := 3
	if !reflect.DeepEqual(sd.Data, expectedData) || sd.Size != expectedSize || index != 4 {
		t.Errorf("Push() = %v, want %v, size = %d, want %d, index = %d, want %d", sd.Data, expectedData, sd.Size, expectedSize, index, 4)
	}
}

func TestPop(t *testing.T) {
	sd := &SliceData[int]{Data: []int{1, 2, 3}, Size: 3}
	ele := sd.Pop()
	expectedData := []int{1, 2}
	if ele != 3 || !reflect.DeepEqual(sd.Data, expectedData) || sd.Size != 2 {
		t.Errorf("Pop() = %d, want %d, resulting slice = %v, want %v, size = %d, want %d", ele, 3, sd.Data, expectedData, sd.Size, 2)
	}
}

func TestUnshift(t *testing.T) {
	sd := &SliceData[int]{Data: []int{3, 4, 5}, Size: 3}
	sliceLen := sd.Unshift(1, 2)
	expectedData := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(sd.Data, expectedData) || sd.Size != 5 || sliceLen != 5 {
		t.Errorf("Unshift() = %v, want %v, size = %d, want %d, index = %d, want %d", sd.Data, expectedData, sd.Size, 5, sliceLen, 5)
	}
}

func TestShift(t *testing.T) {
	sd := &SliceData[int]{Data: []int{1, 2, 3}, Size: 3}
	ele := sd.Shift()
	expectedData := []int{2, 3}
	if ele != 1 || !reflect.DeepEqual(sd.Data, expectedData) || sd.Size != 2 {
		t.Errorf("Shift() = %d, want %d, resulting slice = %v, want %v, size = %d, want %d", ele, 1, sd.Data, expectedData, sd.Size, 2)
	}
}

func TestGetSize(t *testing.T) {
	sd := &SliceData[int]{Size: 5}
	size := sd.GetSize()
	if size != 5 {
		t.Errorf("GetSize() = %d, want %d", size, 5)
	}
}

func TestGetElement(t *testing.T) {
	sd := &SliceData[int]{Data: []int{1, 2, 3, 4, 5}, Size: 5}
	ele := sd.GetElement(2)
	if ele != 3 {
		t.Errorf("GetElement() = %d, want %d", ele, 3)
	}
}

// Add more tests for other methods as necessary
