package pokedex

import (
	"sync"
)

type PokedexEntry struct {
	Name           string
	URL            string
	BaseExperience int
}

type Pokedex struct {
	entries map[string]PokedexEntry
	mutex   sync.RWMutex
}

func New() *Pokedex {
	return &Pokedex{
		entries: make(map[string]PokedexEntry),
	}
}

func (p *Pokedex) Add(name string, entry PokedexEntry) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	p.entries[name] = entry
}

func (p *Pokedex) List() []PokedexEntry {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	pokemon := make([]PokedexEntry, 0, len(p.entries))
	for _, entry := range p.entries {
		pokemon = append(pokemon, entry)
	}
	return pokemon
}
