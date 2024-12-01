package pokedex

import (
	"sync"

	"github.com/jamesfulreader/pokedexcli/internal/models"
)

type Pokedex struct {
	entries map[string]models.Pokemon
	mutex   sync.RWMutex
}

func New() *Pokedex {
	return &Pokedex{
		entries: make(map[string]models.Pokemon),
	}
}

func (p *Pokedex) Add(name string, entry models.Pokemon) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.entries[name] = entry
}

func (p *Pokedex) List() []models.Pokemon {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	pokemon := make([]models.Pokemon, 0, len(p.entries))
	for _, entry := range p.entries {
		pokemon = append(pokemon, entry)
	}
	return pokemon
}

func (p *Pokedex) Inspect(pokemonName string) (models.Pokemon, bool) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	pokemon, exists := p.entries[pokemonName]
	return pokemon, exists
}
