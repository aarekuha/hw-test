package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	if cache.capacity == 0 {
		return false
	}
	item, previouslyExists := cache.items[key]
	if previouslyExists {
		cache.queue.MoveToFront(item)
		item.Value = cacheItem{
			key:   key,
			value: value,
		}
	} else {
		if cache.queue.Len() == cache.capacity {
			last := cache.queue.Back()
			evicted := last.Value.(cacheItem)
			delete(cache.items, evicted.key)
			cache.queue.Remove(last)
		}
		item = cache.queue.PushFront(cacheItem{
			key:   key,
			value: value,
		})
	}
	cache.items[key] = item
	return previouslyExists
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	item, exists := cache.items[key]
	if exists {
		cache.queue.MoveToFront(item)
		return item.Value.(cacheItem).value, true
	}
	return nil, false
}

func (cache *lruCache) Clear() {
	cache.queue = NewList()
	cache.items = make(map[Key]*ListItem, cache.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
