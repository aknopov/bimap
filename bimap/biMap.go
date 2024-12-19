// Package bimap implements bi-directional map similar to Guava -
// see https://guava.dev/releases/14.0/api/docs/com/google/common/collect/BiMap.html
package bimap

// BiMap represents bidirectional map that has unique sets of keys AND values
type BiMap[K comparable, V comparable] struct {
	keyMap   map[K]V
	valueMap map[V]K
}

// Void type for key and values "sets"
type Void struct{}

var Null Void

// Creates a new BiMap. Both Key and Value types should be comparable
func NewBiMap[K comparable, V comparable]() *BiMap[K, V] {
	biMap := BiMap[K, V]{keyMap: make(map[K]V), valueMap: make(map[V]K)}
	return &biMap
}

// Provides bi-map size
func (biMap *BiMap[K, V]) Size() int {
	return len(biMap.keyMap)
}

// Adds or replaces an entry in bi-map
//   - key - the entry key
//   - val - the entry value
// Returns this bi-map
func (biMap *BiMap[K, V]) Put(key K, val V) *BiMap[K, V] {
	if oldValue, ok := biMap.keyMap[key]; ok {
		delete(biMap.valueMap, oldValue)
	}
	biMap.keyMap[key] = val
	biMap.valueMap[val] = key
	return biMap
}

// Gets value by the key
//   - key - the map key
// Returns found value and a flag of success
func (biMap *BiMap[K, V]) GetValue(key K) (V, bool) {
	val, ok := biMap.keyMap[key]
	return val, ok
}

// Gets key by the value
//   - val - value of the matching entry
// Returns found key and a flag of success
func (biMap *BiMap[K, V]) GetKey(val V) (K, bool) {
	key, ok := biMap.valueMap[val]
	return key, ok
}

// Checks if the key is present in the map
func (biMap *BiMap[K, V]) ContainsKey(key K) bool {
	_, ok := biMap.keyMap[key]
	return ok
}

// Checks if value is present in the map
func (biMap *BiMap[K, V]) ContainsValue(value V) bool {
	_, ok := biMap.valueMap[value]
	return ok
}

// Removes entry from bi-map based on a key
//   - biMap - bi-map to update
//   - key - key of the entry bo be removed
func (biMap *BiMap[K, V]) RemoveKey(key K) {
	if val, ok := biMap.keyMap[key]; ok {
		if _, ok = biMap.valueMap[val]; ok {
			delete(biMap.keyMap, key)
			delete(biMap.valueMap, val)
		}
	}
}

// Removes entry from bi-map based on a value
//   - biMap - bi-map to update
//   - val - value of the entry bo be removed
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

// Compares bi-map to the other
//   - biMap - first bi-map to compare
//   - other - second bi-map to compare
func (biMap *BiMap[K, V]) Equals(other *BiMap[K, V]) bool {
	if biMap.Size() != other.Size() {
		return false
	}

	for k, v := range biMap.keyMap {
		if v != biMap.keyMap[k] {
			return false
		}
	}

	return true
}

// Copies all of the mappings from another map to this
//   - biMap - bi-map to copy to
//   - other - bi-map to copy from
func (biMap *BiMap[K, V]) PutAll(other *BiMap[K, V]) *BiMap[K, V] {
	for k, v := range other.keyMap {
		biMap.Put(k, v)
	}
	return biMap
}

//---------------------

// Returns a "set" of bi-map keys
func (biMap *BiMap[K, V]) Keys() map[K]Void {
	ret := make(map[K]Void, len(biMap.keyMap))

	for key := range biMap.keyMap {
		ret[key] = Null
	}
	return ret
}

// Returns a "set" of bi-map values
func (biMap *BiMap[K, V]) Values() map[V]Void {
	ret := make(map[V]Void, len(biMap.valueMap))

	for val := range biMap.valueMap {
		ret[val] = Null
	}
	return ret
}
