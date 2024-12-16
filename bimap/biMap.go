package bimap

//UC See https://guava.dev/releases/14.0/api/docs/com/google/common/collect/BiMap.html
// BiMap represents bidirectional map that has unique sets of keys AND values
type BiMap[K comparable, V comparable] struct {
	keyMap   map[K]V
	valueMap map[V]K
}

func NewBiMap[K comparable, V comparable]() (*BiMap[K, V]) {

	//WIP See https://www.reddit.com/r/golang/comments/17qhp9f/generic_type_as_map_key/
	//WIP https://www.reddit.com/r/golang/comments/zptqqa/generic_constraint_for_maps/
	// What is comparable interface?
	biMap := BiMap[K, V]{keyMap: make(map[K]V), valueMap: make(map[V]K)}
	return &biMap
}

type BiMapOps[K comparable, V comparable] interface {
	Size() (int)
	Put(key K, value V) (*BiMap[K, V], error)
	GetValue(key K) (V)
	GetKey(value V) (K)
	RemoveKey(key K)
	RemoveValue(value V)
}

// Provides bi-map size
func (biMap *BiMap[K, V]) Size() (int) {
	return len(biMap.keyMap)
}

// Adds and entry to the bi-map
func (biMap *BiMap[K, V]) Put(key K, value V) (*BiMap[K, V]) {
	biMap.keyMap[key] = value
	biMap.valueMap[value] = key
	return biMap
}

// Gets value by the key
func (biMap *BiMap[K, V]) GetValue(key K) (V, bool) {
	val, ok := biMap.keyMap[key]
	return val, ok
}

// Gets key by the value
func (biMap *BiMap[K, V]) GetKey(value V) (K, bool) {
	// if key, ok := biMap.valueMap[value]; ok {
	// 	return &key
	// }
	// return nil
	key, ok := biMap.valueMap[value]
	return key, ok
}

// Removes entry from bi-map based on a key
func (biMap *BiMap[K, V])  RemoveKey(key K) {
	if val, ok := biMap.keyMap[key]; ok {
		if _, ok = biMap.valueMap[val]; ok {
			delete(biMap.keyMap, key);
			delete(biMap.valueMap, val);
		}
	}
}

// Removes entry from bi-map based on a value
func (biMap *BiMap[K, V])  RemoveValue(val V) {
	if key, ok := biMap.valueMap[val]; ok {
		if _, ok = biMap.keyMap[key]; ok {
			delete(biMap.keyMap, key);
			delete(biMap.valueMap, val);
		}
	}
}