package cache

import (
	"container/list"
	"fmt"
)

type item struct {
	key   string
	value interface{}
}

type Cache struct {
	dl          *list.List
	capacity    int
	maxCapacity int
	storage     map[string]*list.Element
}

func (c *Cache) Set(key string, value interface{}) {
	e, ok := c.storage[key]
	if ok {
		n := e.Value.(item)
		n.value = value
		e.Value = n
		c.dl.MoveToBack(e)
		return
	}

	if c.capacity >= c.maxCapacity {
		e := c.dl.Front()
		dk := e.Value.(item).key
		c.dl.Remove(e)
		delete(c.storage, dk)
		c.capacity--
	}

	n := item{key: key, value: value}
	ne := c.dl.PushBack(n)
	c.storage[key] = ne
	c.capacity++
}

func (c *Cache) Get(key string) interface{} {
	v, ok := c.storage[key]
	if ok {
		c.dl.MoveToBack(v)
		return v.Value.(item).value
	}
	return nil
}

func (c *Cache) Print() {
	for k, v := range c.storage {
		fmt.Printf("storage -> key: %v, value: %v\n", k, v.Value.(item))
	}
	fmt.Println()
}

func (c *Cache) Printlist() {
	for e := c.dl.Front(); e != nil; e = e.Next() {
		fmt.Printf("%+v", e.Value.(item))
	}
	fmt.Println()
}

func New(cap int) *Cache {
	return &Cache{
		dl:          list.New(),
		storage:     make(map[string]*list.Element),
		capacity:    0,
		maxCapacity: cap,
	}
}
