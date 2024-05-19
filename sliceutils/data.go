package sliceutils

// SliceData T comparable cant be a reference type, for example, slice, map, channel
type SliceData[T comparable] struct {
	Data []T
	Size int
}

func (sd *SliceData[T]) Push(elements ...T) (index int) {
	sd.Data = append(sd.Data, elements...)
	sd.Size += len(elements)
	index = len(sd.Data) + 1
	return index
}

func (sd *SliceData[T]) Pop() (ele T) {
	if sd.Size == 0 {
		return
	}
	ele = sd.Data[sd.Size-1]
	sd.Data = sd.Data[:sd.Size-1]
	sd.Size--
	return ele
}

func (sd *SliceData[T]) Unshift(elements ...T) (sliceLen int) {
	sd.Data = append(elements, sd.Data...)
	sd.Size += len(elements)
	sliceLen = len(sd.Data)
	return sliceLen
}

func (sd *SliceData[T]) Shift() (ele T) {
	if sd.Size == 0 {
		return
	}
	ele = sd.Data[0]
	sd.Data = sd.Data[1:]
	sd.Size--
	return ele
}

func (sd *SliceData[T]) Slice(startIndex int, endIndex int) []T {
	return sd.Data[startIndex:endIndex]
}

func (sd *SliceData[T]) Splice(startIndex int, deleteCount int, elements ...T) []T {
	if deleteCount < 0 {
		deleteCount = 0
	}

	if deleteCount == 0 {
		// add elements
		sd.Data = append(sd.Data[:startIndex], append(elements, sd.Data[startIndex:]...)...)
		return sd.Data
	}

	// delete elements
	sd.Data = append(sd.Data[:startIndex], sd.Data[startIndex+deleteCount:]...)
	return sd.Data
}

func (sd *SliceData[T]) Reverse() (reversed []T) {
	for i := sd.Size - 1; i >= 0; i-- {
		reversed = append(reversed, sd.Data[i])
	}
	return reversed
}

func (sd *SliceData[T]) Concat(slice []T) (concatenated []T) {
	concatenated = append(sd.Data, slice...)
	return concatenated
}

func (sd *SliceData[T]) Includes(ele T) bool {
	for _, v := range sd.Data {
		if v == ele {
			return true
		}
	}
	return false
}

func (sd *SliceData[T]) GetSize() int {
	return sd.Size
}

func (sd *SliceData[T]) GetSlice() []T {
	return sd.Data
}

func (sd *SliceData[T]) GetElement(index int) (ele T) {
	return sd.Data[index]
}

func (sd *SliceData[T]) ForEach(callback func(index int, ele T, slice []T)) {
	for i, v := range sd.Data {
		callback(i, v, sd.Data)
	}
}

func (sd *SliceData[T]) Map(callback func(index int, ele T, slice []T) T) {
	for i, v := range sd.Data {
		sd.Data[i] = callback(i, v, sd.Data)
	}
}
