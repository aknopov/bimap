// Package bimap implements bi-directional map similar to Guava -
// see https://guava.dev/releases/14.0/api/docs/com/google/common/collect/BiMap.html
package bimap

import "reflect"

// BiMap represents bidirectional map that has unique sets of keys AND values
type BiMap[K comparable, V comparable] struct {
	keys   []K
	vals   []V
	keyIdx map[K]int
	valIdx map[V]int
	noKey  K
	noVal  V
}

// Void type for key and values "sets"
type Void struct{}

var Null Void

// Creates a new zero-sized BiMap
func NewBiMap[K comparable, V comparable]() *BiMap[K, V] {
	return NewBiMapEx[K, V](0)
}

// Creates a new BiMap
//   - capacity - initial capacity
func NewBiMapEx[K comparable, V comparable](capacity int) *BiMap[K, V] {
	keys := make([]K, 0, capacity)
	vals := make([]V, 0, capacity)
	keyIdx := make(map[K]int, capacity)
	valIdx := make(map[V]int, capacity)
	return &BiMap[K, V]{keys: keys, vals: vals, keyIdx: keyIdx, valIdx: valIdx}
}

// Provides bi-map size
func (biMap *BiMap[K, V]) Size() int {
	return len(biMap.keys)
}

// Adds or replaces an entry in bi-map
//   - key - the entry key
//   - val - the entry value
//
// Returns this bi-map
func (biMap *BiMap[K, V]) Put(key K, val V) *BiMap[K, V] {
	if i, ok := biMap.keyIdx[key]; ok {
		oldVal := biMap.vals[i]
		delete(biMap.valIdx, oldVal)
		biMap.keys[i] = key
		biMap.vals[i] = val
		biMap.valIdx[val] = i
	} else {
		biMap.keyIdx[key] = len(biMap.keys)
		biMap.valIdx[val] = len(biMap.vals)
		biMap.keys = append(biMap.keys, key)
		biMap.vals = append(biMap.vals, val)
	}
	return biMap
}

// Gets value by the key
//   - key - the map key
//
// Returns found value and a flag of success
func (biMap *BiMap[K, V]) GetValue(key K) (V, bool) {
	if i, ok := biMap.keyIdx[key]; ok {
		return biMap.vals[i], true
	}
	return biMap.noVal, false
}

// Gets key by the value
//   - val - value of the matching entry
//
// Returns found key and a flag of success
func (biMap *BiMap[K, V]) GetKey(val V) (K, bool) {
	if i, ok := biMap.valIdx[val]; ok {
		return biMap.keys[i], true
	}
	return biMap.noKey, false
}

// Checks if the key is present in the map
func (biMap *BiMap[K, V]) ContainsKey(key K) bool {
	_, ok := biMap.keyIdx[key]
	return ok
}

// Checks if value is present in the map
func (biMap *BiMap[K, V]) ContainsValue(value V) bool {
	_, ok := biMap.valIdx[value]
	return ok
}

// Removes entry from bi-map based on a key
//   - biMap - bi-map to update
//   - key - key of the entry bo be removed
func (biMap *BiMap[K, V]) RemoveKey(key K) {
	if i, ok := biMap.keyIdx[key]; ok {
		val := biMap.vals[i]
		biMap.removeEntry(key, val, i)
	}
}

// Removes entry from bi-map based on a value
//   - biMap - bi-map to update
//   - val - value of the entry bo be removed
func (biMap *BiMap[K, V]) RemoveValue(val V) {
	if i, ok := biMap.valIdx[val]; ok {
		key := biMap.keys[i]
		biMap.removeEntry(key, val, i)
	}
}

func (biMap *BiMap[K, V]) removeEntry(key K, val V, i int) {
	newLen := len(biMap.keys) - 1
	delete(biMap.keyIdx, key)
	delete(biMap.valIdx, val)
	biMap.keys = append(biMap.keys[:i], biMap.keys[i+1:]...)
	biMap.vals = append(biMap.vals[:i], biMap.vals[i+1:]...)
	for j := i; j < newLen; j++ {
		biMap.keyIdx[biMap.keys[j]] = j
		biMap.valIdx[biMap.vals[j]] = j
	}
}

// Creates "inverse" copy of the bitmap
func (biMap *BiMap[K, V]) Inverse() *BiMap[V, K] {
	size := biMap.Size()
	invMap := BiMap[V,K]{}
	invMap.keys = append(invMap.keys, biMap.vals...)
	invMap.vals = append(invMap.vals, biMap.keys...)
	invMap.keyIdx = make(map[V]int, size)
	invMap.valIdx = make(map[K]int, size)
	for i := 0; i < size; i++ {
		invMap.keyIdx[invMap.keys[i]] = i
		invMap.valIdx[invMap.vals[i]] = i
	}
	return &invMap
}

// Compares bi-map to the other
//   - biMap - first bi-map to compare
//   - other - second bi-map to compare
func (biMap *BiMap[K, V]) Equals(other *BiMap[K, V]) bool {
	return  reflect.DeepEqual(biMap.keys, other.keys) && reflect.DeepEqual(biMap.vals, other.vals)
}

// Copies all of the mappings from another map to this
//   - biMap - bi-map to copy to
//   - other - bi-map to copy from
func (biMap *BiMap[K, V]) PutAll(other *BiMap[K, V]) *BiMap[K, V] {
	for i := range other.keys {
		biMap.Put(other.keys[i], other.vals[i])
	}
	return biMap
}

// Returns a slice of bi-map keys
func (biMap *BiMap[K, V]) Keys() []K {
	return biMap.keys
}

// Returns a slice of bi-map values
func (biMap *BiMap[K, V]) Values() []V {
	return biMap.vals
}
