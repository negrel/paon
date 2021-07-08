package id

import "sync"

// StringMap define a map that use IDs as keys and string as values.
type StringMap interface {
	Set(p ID, name string)
	Get(p ID) string
	Del(p ID)
}

var _ StringMap = &idMap{}

type idMap struct {
	sync.RWMutex
	m map[ID]string
}

// NewMap returns a new empty StringMap
func NewMap() StringMap {
	return &idMap{
		m: make(map[ID]string),
	}
}

func (im *idMap) Set(p ID, name string) {
	im.Lock()
	defer im.Unlock()
	im.m[p] = name
}

func (im *idMap) Get(p ID) string {
	im.RLock()
	defer im.RUnlock()
	return im.m[p]
}

func (im *idMap) Del(p ID) {
	im.Lock()
	defer im.Unlock()
	delete(im.m, p)
}
