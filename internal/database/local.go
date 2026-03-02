package database

import "sync"

type LocalStorage struct {
	mx sync.RWMutex
	short map[int64]string
	full  map[string]int64
}

func (l *LocalStorage) Init() {
	l.short = make(map[int64]string)
	l.full = make(map[string]int64)
}

func (l *LocalStorage) Save(shortLink int64, fullLink string) {
	l.mx.Lock()
	l.short[shortLink] = fullLink
	l.full[fullLink] = shortLink
	l.mx.Unlock()
}

func (l *LocalStorage) FindFull(shortLink int64) (string, bool) {
	l.mx.RLock()
	value, ok := l.short[shortLink]
	l.mx.RUnlock()
	return value, ok
}

func (l *LocalStorage) FindShort(fullLink string) (int64, bool) {
	l.mx.RLock()
	value, ok := l.full[fullLink]
	l.mx.RUnlock()
	return value, ok
}

func (l *LocalStorage) Close() {
	l.short = nil
	l.full = nil
}