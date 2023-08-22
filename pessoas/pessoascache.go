package pessoas

import (
	"sync"
	"time"
)

type PessoaCache struct {
	cache   map[string]*Pessoa
	timeout time.Duration
	mutex   sync.RWMutex
}

var (
	pessoasCache *PessoaCache
)

func init() {
	pessoasCache = &PessoaCache{
		cache:   make(map[string]*Pessoa),
		timeout: 5 * time.Minute,
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

	go func(id string, timeout time.Duration) {
		time.Sleep(timeout)
		removePessoaFromCache(id)
	}(p.Id, pessoasCache.timeout)
}

func removePessoaFromCache(id string) {
	pessoasCache.mutex.Lock()
	defer pessoasCache.mutex.Unlock()

	delete(pessoasCache.cache, id)
}
