package pessoas

import (
	"sync"
)

type PessoaCache struct {
	cache map[string]*Pessoa
	mutex sync.RWMutex
}

var (
	pessoasCache *PessoaCache
)

func init() {
	pessoasCache = &PessoaCache{
		cache: make(map[string]*Pessoa),
	}
}

func getPessoaFromCache(id string) *Pessoa {
	pessoasCache.mutex.RLock()
	defer pessoasCache.mutex.RUnlock()

	p, exists := pessoasCache.cache[id]
	if exists {
		return p
	}
	return nil
}

func addPessoaToCache(p *Pessoa) {
	pessoasCache.mutex.Lock()
	defer pessoasCache.mutex.Unlock()

	pessoasCache.cache[p.Id] = p
}
