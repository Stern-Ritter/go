package main

import "sync"

type Counter struct {
	mu    sync.RWMutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Add(delta int) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += delta
	return c.value
}

func (c *Counter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}
