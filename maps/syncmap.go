package maps

import "sync"

type TypedSyncMap[K any, V any] sync.Map

func (m *TypedSyncMap[K, V]) Store(key K, value V) {
	(*sync.Map)(m).Store(key, value)
}

func (m *TypedSyncMap[K, V]) Load(key K) (V, bool) {
	r, ok := (*sync.Map)(m).Load(key)
	if ok {
		return r.(V), ok
	}
	var zero V
	return zero, ok
}

func (m *TypedSyncMap[K, V]) LoadOrStore(key K, value V) (V, bool) {
	r, loaded := (*sync.Map)(m).LoadOrStore(key, value)
	return r.(V), loaded
}

func (m *TypedSyncMap[K, V]) LoadAndDelete(key K) (V, bool) {
	r, ok := (*sync.Map)(m).LoadAndDelete(key)
	if ok {
		return r.(V), ok
	}
	var zero V
	return zero, ok
}

func (m *TypedSyncMap[K, V]) Delete(key K) {
	(*sync.Map)(m).Delete(key)
}

func (m *TypedSyncMap[K, V]) Range(f func(key K, value V) bool) {
	ff := func(key any, value any) bool {
		return f(key.(K), value.(V))
	}
	(*sync.Map)(m).Range(ff)
}

func NewTypedSyncMap[K any, V any]() *TypedSyncMap[K, V] {
	return &TypedSyncMap[K, V]{}
}
