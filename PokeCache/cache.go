package PokeCache

import (
	"sync"
	"time"
	"github.com/Sirkle23/PokedexGo/PokedexAPI"
)

type entry struct {
	createdAt time.Time
	data      []byte
}

type Cache struct {
	cacheEntry map[string]entry
	mu	sync.Mutex
	interval time.Duration
	CaughtPokemon map[string]PokedexAPI.Pokemon_strct
}

func NewCache(interval time.Duration) *Cache{
	c := &Cache{
		cacheEntry: make(map[string]entry),
		CaughtPokemon: make(map[string]PokedexAPI.Pokemon_strct),
		interval:   interval,
	}

	go c.reapLoop()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	newEntry := entry{
		createdAt: time.Now(),
		data:      val,
	}

	c.cacheEntry[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.cacheEntry[key]
	if !exists {
		return nil, false
	}

	return entry.data, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.cacheEntry {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.cacheEntry, key)
			}
		}
		c.mu.Unlock()
	}
}

func (c *Cache) AddPokemon(name string, pokemon PokedexAPI.Pokemon_strct) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.CaughtPokemon[name] = pokemon
}

func (c *Cache) GetPokemon(name string) (PokedexAPI.Pokemon_strct, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	pokemon, exists := c.CaughtPokemon[name]
	if !exists {
		return PokedexAPI.Pokemon_strct{}, false
	}

	return pokemon, true
}