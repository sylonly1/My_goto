package gotov1

import (
	"log"
	"sync"
)

type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
}

func NewURLStore() *URLStore {
	return &URLStore{urls: map[string]string{}}
}

func (u *URLStore) Get(key string) string {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.urls[key]
}

func (u *URLStore) Set(key, url string) bool {
	u.mu.Lock()
	defer u.mu.Unlock()

	if _, ok := u.urls[key];ok{
		return false
	}

	u.urls[key] = url
	return true
}

func (u *URLStore) Count() int {
	u.mu.RLock()
	defer u.mu.RUnlock()

	return len(u.urls)
}

func (u *URLStore) Put(url string) string {
	for i := 0; i < 100; i++ {
		key := genKey(u.Count())
		if u.Set(key, url) {
			return key
		}
	}

	log.Printf("Set url: %s Error", url)
	return ""
}