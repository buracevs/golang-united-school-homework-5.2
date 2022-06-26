package cache

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	actual := NewCache()

	if len(actual.values) != 0 {
		t.Error("length should be 0")
	}
}

func TestPutAddsNewItem(t *testing.T) {
	actual := NewCache()

	actual.Put("tst", "data")
	keys := actual.Keys()
	if len(keys) < 1 {
		t.Error("Items should be found")
	}
}

func TestKeysReturnsCorrectKey(t *testing.T) {
	actual := NewCache()

	actual.Put("tst", "data")
	keys := actual.Keys()
	if keys[0] != "tst" {
		t.Errorf("Items should be found expected tst but was : %v", keys[0])
	}
}

func TestPutOverridesValue(t *testing.T) {
	actual := NewCache()

	actual.Put("tst", "data")
	actual.Put("tst", "data 111")
	value, _ := actual.Get("tst")
	if value != "data 111" {
		t.Error("Data should be changed")
	}
}

func TestKeysReturnsData(t *testing.T) {
	actual := NewCache()
	nowPlus := time.Now().Add(time.Minute * 2)
	actual.PutTill("tst", "data", nowPlus)
	keys := actual.Keys()
	if len(keys) < 1 {
		t.Error("Items should be found")
	}
}

func TestGetReturnsFalseAsResult(t *testing.T) {
	actual := NewCache()

	_, found := actual.Get("tst")
	if found {
		t.Error("No items should be found")
	}
}

func TestGetHasValueNoDeadline(t *testing.T) {
	actual := NewCache()
	actual.Put("tst", "data")
	_, found := actual.Get("tst")
	if !found {
		t.Error("Items should be found")
	}
}

func TestGetHasValueHasDeadline(t *testing.T) {
	actual := NewCache()
	nowPlus := time.Now().Add(time.Minute * 2)
	actual.PutTill("tst", "data", nowPlus)
	_, found := actual.Get("tst")
	if !found {
		t.Error("Items should be found")
	}
}
