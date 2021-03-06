// Package _test provides utilities for cache testing
package _test

import (
	"sync"
	"testing"

	"github.com/pierrre/imageserver"
	imageserver_cache "github.com/pierrre/imageserver/cache"
	"github.com/pierrre/imageserver/testdata"
)

const (
	// KeyValid is a valid cache key (with content)
	KeyValid = "test"
	// KeyMiss is an invalid cache key (without content)
	KeyMiss = "unknown"
)

// CacheTestGetSet is a helper to test cache Get()/Set()
func CacheTestGetSet(t *testing.T, cache imageserver_cache.Cache, image *imageserver.Image) {
	err := cache.Set(KeyValid, image, imageserver.Params{})
	if err != nil {
		t.Fatal(err)
	}

	newImage, err := cache.Get(KeyValid, imageserver.Params{})
	if err != nil {
		t.Fatal(err)
	}

	if !imageserver.ImageEqual(newImage, image) {
		t.Fatal("image not equals")
	}
}

// CacheTestGetSetAllImages is a helper to test cache Get()/Set() with all images from test data
func CacheTestGetSetAllImages(t *testing.T, cache imageserver_cache.Cache) {
	for _, image := range testdata.Images {
		CacheTestGetSet(t, cache, image)
	}
}

// CacheTestGetErrorMiss is a helper to test cache Get() with a "cache miss" error
func CacheTestGetErrorMiss(t *testing.T, cache imageserver_cache.Cache) {
	_, err := cache.Get(KeyMiss, imageserver.Params{})
	if err == nil {
		t.Fatal("no error")
	}
	if _, ok := err.(*imageserver_cache.MissError); !ok {
		t.Fatal("invalid error type")
	}
}

// MapCache is a simple Cache (it wraps a map) for tests
type MapCache struct {
	mutex sync.RWMutex
	data  map[string]*imageserver.Image
}

// NewMapCache creates a new CacheMap
func NewMapCache() *MapCache {
	return &MapCache{
		data: make(map[string]*imageserver.Image),
	}
}

// Get gets an Image from the MapCache
func (cache *MapCache) Get(key string, params imageserver.Params) (*imageserver.Image, error) {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	image, ok := cache.data[key]
	if !ok {
		return nil, &imageserver_cache.MissError{Key: key}
	}

	return image, nil
}

// Set sets an Image to the MapCache
func (cache *MapCache) Set(key string, image *imageserver.Image, params imageserver.Params) error {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	cache.data[key] = image

	return nil
}
