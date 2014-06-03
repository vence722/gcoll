// gcoll v1.0
// @description gcoll is a go collection library which you can use like in Java
// @authors     Vence Lin(vence722@gmail.com)

package maps

import (
	"fmt"
)

// The HashMap struct
type HashMap struct {
	hmap map[interface{}]interface{}
}

// Return a new HashMap
func NewHashMap() *HashMap {
	hmap := make(map[interface{}]interface{})
	return &HashMap{hmap}
}

// Return the size of the map
func (this *HashMap) Size() int {
	return len(this.hmap)
}

// Return the map containing elements or not
func (this *HashMap) IsEmpty() bool {
	return len(this.hmap) == 0
}

// Return the map conaining specified key or not
func (this *HashMap) ContainsKey(key interface{}) bool {
	for k, _ := range this.hmap {
		if k == key {
			return true
		}
	}
	return false
}

// Return the map conaining specified value or not
func (this *HashMap) ContainsValue(value interface{}) bool {
	for _, v := range this.hmap {
		if v == value {
			return true
		}
	}
	return false
}

// Return the element with the specified key in this map
func (this *HashMap) Get(key interface{}) interface{} {
	return this.hmap[key]
}

// Add new key value pair to this map
func (this *HashMap) Put(key, value interface{}) bool {
	this.hmap[key] = value
	return true
}

// Remove value with specified key from this map
func (this *HashMap) Remove(key interface{}) bool {
	delete(this.hmap, key)
	return true
}

// Add another map to this map
func (this *HashMap) PutAll(amap Map) bool {
	entries := amap.Entries()
	for i := 0; i < len(entries); i++ {
		this.Put(entries[i].Key, entries[i].Value)
	}
	return true
}

// Remove all elements from this map
func (this *HashMap) Clear() {
	for k, _ := range this.hmap {
		delete(this.hmap, k)
	}
}

// Views
// Return a list containing all the keys in the map
func (this *HashMap) Keys() []interface{} {
	keys := make([]interface{}, INIT_LEN, INIT_CAP)
	for k, _ := range this.hmap {
		keys = append(keys, k)
	}
	return keys
}

// Return a list containing all the values in the map
func (this *HashMap) Values() []interface{} {
	values := make([]interface{}, INIT_LEN, INIT_CAP)
	for _, v := range this.hmap {
		values = append(values, v)
	}
	return values
}

// Return a list containing all the entries in the map
func (this *HashMap) Entries() []MapEntry {
	entries := make([]MapEntry, INIT_LEN, INIT_CAP)
	for k, v := range this.hmap {
		entries = append(entries, MapEntry{k, v})
	}
	return entries
}

// Return the string that describes the contains of this map
func (this *HashMap) String() string {
	return fmt.Sprint(this.hmap)
}
