package cache

import "time"

type Cache struct {
}

func NewCache() Cache {
	return Cache{}
}

func (receiver) Get(key string) (string, bool) {

}

func (receiver) Put(key, value string) {
}

func (receiver) Keys() []string {
}

func (receiver) PutTill(key, value string, deadline time.Time) {
}
