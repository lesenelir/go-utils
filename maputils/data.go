package maputils

type MapData[K comparable, V any] struct {
	Map  map[K]V
	Size int
}

func (md *MapData[K, V]) Set(key K, value V) {
	md.Map[key] = value
	md.Size += 1
}

func (md *MapData[K, V]) Get(key K) V {
	return md.Map[key]
}

func (md *MapData[K, V]) Delete(key K) {
	delete(md.Map, key)
	md.Size -= 1
}

func (md *MapData[K, V]) Has(key K) bool {
	_, ok := md.Map[key]
	return ok
}

func (md *MapData[K, V]) Clear() {
	md.Map = make(map[K]V)
	md.Size = 0
}
