package maputils

import (
	"reflect"
	"testing"
)

func TestMapData_SetAndGet(t *testing.T) {
	md := &MapData[string, int]{
		Map:  make(map[string]int),
		Size: 0,
	}

	md.Set("a", 1)
	md.Set("b", 2)
	md.Set("c", 3)

	expectedData := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	expectedSize := 3
	expectedEle := md.Get("b")

	if !reflect.DeepEqual(md.Map, expectedData) || md.Size != expectedSize || expectedEle != 2 {
		t.Errorf("Set or Get failed: got %v, want %v, size got %d, want %d", md.Map, expectedData, md.Size, expectedSize)
	}
}

func TestMapData_DeleteAndHas(t *testing.T) {
	md := &MapData[string, int]{
		Map:  make(map[string]int),
		Size: 0,
	}

	md.Set("a", 1)
	md.Set("b", 2)
	md.Set("c", 3)

	md.Delete("b")
	exists := md.Has("a")

	expectedData := map[string]int{
		"a": 1,
		"c": 3,
	}
	expectedSize := 2

	if _, ok := md.Map["b"]; ok || !reflect.DeepEqual(md.Map, expectedData) || md.Size != expectedSize || exists != true {
		t.Errorf("Delete failed: map got %v, want %v, size got %d, want %d", md.Map, expectedData, md.Size, expectedSize)
	}
}

func TestMapData_Clear(t *testing.T) {
	md := &MapData[string, int]{
		Map:  make(map[string]int),
		Size: 0,
	}

	md.Set("a", 1)
	md.Set("b", 2)
	md.Set("c", 3)

	md.Clear()

	expectedMap := map[string]int{}
	expectedSize := 0

	if len(md.Map) != 0 || !reflect.DeepEqual(md.Map, expectedMap) || md.Size != expectedSize {
		t.Errorf("Clear failed: map got %v, want %v, size got %d, want %d", md.Map, expectedMap, md.Size, expectedSize)
	}
}
