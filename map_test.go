package async_tools

import (
	"testing"
)

type testData struct {
	Key   string
	Value int
}

func TestSyncMap(t *testing.T) {
	cm := SyncMap[int]{}

	testData := []testData{
		{
			Key:   "1",
			Value: 1,
		},
		{
			Key:   "2",
			Value: 2,
		}, {
			Key:   "3",
			Value: 3,
		}, {
			Key:   "4",
			Value: 4,
		},
	}

	for _, v := range testData {
		cm.Store(v.Key, v.Value)
	}

	for _, v := range testData {
		value, ok := cm.Load(v.Key)
		if !ok {
			t.Fatal("load 1 error")
		}
		if *value != v.Value {
			t.Fatal("load 1 error")
		}
	}

	cm.Range(func(key string, value *int) bool {
		for _, v := range testData {
			if key == v.Key {
				if v.Value != *value {
					t.Fatal("load 1 error")
				}
			}
		}

		return true
	})
}

func TestRWMap(t *testing.T) {
	cm := RWMap[int]{}

	testData := []testData{
		{
			Key:   "1",
			Value: 1,
		},
		{
			Key:   "2",
			Value: 2,
		}, {
			Key:   "3",
			Value: 3,
		}, {
			Key:   "4",
			Value: 4,
		},
	}

	for _, v := range testData {
		cm.Store(v.Key, v.Value)
	}

	for _, v := range testData {
		value, ok := cm.Load(v.Key)
		if !ok {
			t.Fatal("load 1 error")
		}
		if *value != v.Value {
			t.Fatal("load 1 error")
		}
	}

	cm.Range(func(key string, value *int) bool {
		for _, v := range testData {
			if key == v.Key {
				if v.Value != *value {
					t.Fatal("load 1 error")
				}
			}
		}

		return true
	})
}
