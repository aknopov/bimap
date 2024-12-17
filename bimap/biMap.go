// Package bimap implements bi-directional map similar to Guava -
// see https://guava.dev/releases/14.0/api/docs/com/google/common/collect/BiMap.html
package bimap

// BiMap represents bidirectional map that has unique sets of keys AND values
type BiMap[K comparable, V comparable] struct {
	keyMap   map[K]V
	valueMap map[V]K
}

// Creates a new BiMap. Both Key and Value types should be comparable
func NewBiMap[K comparable, V comparable]() *BiMap[K, V] {
	biMap := BiMap[K, V]{keyMap: make(map[K]V), valueMap: make(map[V]K)}
	return &biMap
}

// Provides bi-map size
func (biMap *BiMap[K, V]) Size() int {
	return len(biMap.keyMap)
}

// Adds and entry to the bi-map
func (biMap *BiMap[K, V]) Put(key K, value V) *BiMap[K, V] {
	biMap.keyMap[key] = value
	biMap.valueMap[value] = key
	return biMap
}

// Gets value by the key
func (biMap *BiMap[K, V]) GetValue(key K) (V, bool) {
	val, ok := biMap.keyMap[key]
	return val, ok
}

// Checks if key is present in the map
func (biMap *BiMap[K, V]) ContainsKey(key K) bool {
	_, ok := biMap.keyMap[key]
	return ok
}

// Checks if value is present in the map
func (biMap *BiMap[K, V]) ContainsValue(value V) bool {
	_, ok := biMap.valueMap[value]
	return ok
}

// Gets key by the value
func (biMap *BiMap[K, V]) GetKey(value V) (K, bool) {
	key, ok := biMap.valueMap[value]
	return key, ok
}

// Removes entry from bi-map based on a key
func (biMap *BiMap[K, V]) RemoveKey(key K) {
	if val, ok := biMap.keyMap[key]; ok {
		if _, ok = biMap.valueMap[val]; ok {
			delete(biMap.keyMap, key)
			delete(biMap.valueMap, val)
		}
	}
}

// Removes entry from bi-map based on a value
func (biMap *BiMap[K, V]) RemoveValue(val V) {
	if key, ok := biMap.valueMap[val]; ok {
		if _, ok = biMap.keyMap[key]; ok {
			delete(biMap.keyMap, key)
			delete(biMap.valueMap, val)
		}
	}
}

// Creates "inverse" copy of the bitmap
func (biMap *BiMap[K, V]) Inverse() *BiMap[V, K] {
	ret := NewBiMap[V, K]()
	for key, val := range biMap.keyMap {
		ret.keyMap[val] = key
		ret.valueMap[key] = val
	}
	return ret
}
