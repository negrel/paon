package idmap

import "sync"

type Map interface {
	Set(p int32, name string)
	Get(p int32) string
	Del(p int32)
}

var _ Map = &idMap{}

type idMap struct {
	sync.RWMutex
	m map[int32]string
}

func New(len int) Map {
	return &idMap{
		m: make(map[int32]string, len),
	}
}

func (im *idMap) Set(p int32, name string) {
	im.Lock()
	defer im.Unlock()
	im.m[p] = name
}

func (im *idMap) Get(p int32) string {
	im.RLock()
	defer im.RUnlock()
	return im.m[p]
}

func (im *idMap) Del(p int32) {
	im.Lock()
	defer im.Unlock()
	delete(im.m, p)
}
