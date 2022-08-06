package async_tools

import "sync"

// internal map

type internalMap[T base] struct {
	dirty map[string]*T
}

func (s *internalMap[T]) init() {
	if s.dirty == nil {
		s.dirty = map[string]*T{}
	}
}

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (s *internalMap[T]) Load(key string) (value *T, ok bool) {
	s.init()

	t, ok := s.dirty[key]
	return t, ok
}

// Store sets the value for a key.
func (s *internalMap[T]) Store(key string, value T) {
	s.init()

	s.dirty[key] = &value
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (s *internalMap[T]) LoadAndDelete(key string) (value *T, loaded bool) {
	s.init()

	t, ok := s.dirty[key]
	delete(s.dirty, key)

	if !ok {
		return nil, false
	}

	return t, ok
}

// Delete deletes the value for a key.
func (s *internalMap[T]) Delete(key string) {
	s.init()

	delete(s.dirty, key)
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (s *internalMap[T]) LoadOrStore(key string, value T) (actual any, loaded bool) {
	s.init()

	t, ok := s.dirty[key]
	if ok {
		return t, ok
	}

	s.dirty[key] = &value

	return &value, false
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently (including by f), Range may reflect any
// mapping for that key from any point during the Range call. Range does not
// block other methods on the receiver; even f itself may call any method on m.
//
// Range may be O(N) with the number of elements in the map even if f returns
// false after a constant number of calls.
func (s *internalMap[T]) Range(f func(key string, value *T) bool) {
	s.init()

	for k, v := range s.dirty {
		if !f(k, v) {
			break
		}
	}
}

type SyncMap[T base] struct {
	mu    sync.Mutex
	dirty internalMap[T]
}

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (s *SyncMap[T]) Load(key string) (value *T, ok bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.dirty.Load(key)
}

// Store sets the value for a key.
func (s *SyncMap[T]) Store(key string, value T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.dirty.Store(key, value)
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (s *SyncMap[T]) LoadAndDelete(key string) (value *T, loaded bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.dirty.LoadAndDelete(key)
}

// Delete deletes the value for a key.
func (s *SyncMap[T]) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.dirty.Delete(key)
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (s *SyncMap[T]) LoadOrStore(key string, value T) (actual any, loaded bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.dirty.LoadOrStore(key, value)
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently (including by f), Range may reflect any
// mapping for that key from any point during the Range call. Range does not
// block other methods on the receiver; even f itself may call any method on m.
//
// Range may be O(N) with the number of elements in the map even if f returns
// false after a constant number of calls.
func (s *SyncMap[T]) Range(f func(key string, value *T) bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.dirty.Range(f)
}

type RWMap[T base] struct {
	mu    sync.RWMutex
	dirty internalMap[T]
}

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (s *RWMap[T]) Load(key string) (value *T, ok bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.dirty.Load(key)
}

// Store sets the value for a key.
func (s *RWMap[T]) Store(key string, value T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.dirty.Store(key, value)
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (s *RWMap[T]) LoadAndDelete(key string) (value *T, loaded bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.dirty.LoadAndDelete(key)
}

// Delete deletes the value for a key.
func (s *RWMap[T]) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.dirty.Delete(key)
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (s *RWMap[T]) LoadOrStore(key string, value T) (actual any, loaded bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.dirty.LoadOrStore(key, value)
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently (including by f), Range may reflect any
// mapping for that key from any point during the Range call. Range does not
// block other methods on the receiver; even f itself may call any method on m.
//
// Range may be O(N) with the number of elements in the map even if f returns
// false after a constant number of calls.
func (s *RWMap[T]) Range(f func(key string, value *T) bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	s.dirty.Range(f)
}
