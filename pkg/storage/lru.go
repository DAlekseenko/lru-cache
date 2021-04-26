package storage

import (
	"container/list"
	"sync"
)

type LRUCache interface {
	// Add Добавляет новое значение с ключом в кеш (с наивысшим приоритетом), возвращает true, если все прошло успешно
	// В случае дублирования ключа вернуть false
	// В случае превышения размера - вытесняется наименее приоритетный элемент
	Add(key, value string) bool

	// Get Возвращает значение под ключом и флаг его наличия в кеше
	// В случае наличия в кеше элемента повышает его приоритет
	Get(key string) (value string, ok bool)

	// Remove Удаляет элемент из кеша, в случае успеха возвращает true, в случае отсутствия элемента - false
	Remove(key string) (ok bool)
}

type Item struct {
	key   string
	value string
}

type LRU struct {
	capacity int
	storage  map[string]*list.Element
	queue    *list.List
	mu       *sync.RWMutex
}

func (c *LRU) Add(key, value string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, exist := c.storage[key]; exist {
		return false
	}

	if c.queue.Len() == c.capacity {
		c.purge()
	}

	val := &Item{
		key:   key,
		value: value,
	}

	item := c.queue.PushFront(val)
	c.storage[key] = item
	return true
}

func (c *LRU) Get(key string) (value string, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if item, exist := c.storage[key]; exist {
		c.queue.MoveToFront(item)
		return item.Value.(*Item).value, exist
	}
	return "", false
}

func (c *LRU) Remove(key string) (ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if item, exist := c.storage[key]; exist {
		c.queue.Remove(item)
		delete(c.storage, key)
		return true
	}

	return false
}

func (c *LRU) purge() {
	if element := c.queue.Back(); element != nil {
		item := c.queue.Remove(element).(*Item)
		delete(c.storage, item.key)
	}
}

func NewLRUCache(capacity int) *LRU {

	return &LRU{
		capacity: capacity,
		storage:  make(map[string]*list.Element, capacity),
		queue:    list.New(),
		mu:       &sync.RWMutex{},
	}
}
