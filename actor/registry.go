package actor

import "sync"

type registry struct {
	locker sync.RWMutex
	lookup map[string]Processer
}

func newRegistry() *registry {
	return &registry{
		lookup: make(map[string]Processer),
	}
}

func (r *registry) Add(id string, proc Processer) bool {
	r.locker.Lock()
	defer r.locker.Unlock()
	if _, ok := r.lookup[id]; ok {
		return false
	}
	r.lookup[id] = proc
	return true
}

func (r *registry) Remove(pid *PID) {
	r.locker.Lock()
	defer r.locker.Unlock()
	delete(r.lookup, pid.ID)
}

func (r *registry) Get(pid *PID) Processer {
	if pid == nil {
		return nil
	}
	r.locker.RLock()
	defer r.locker.RUnlock()
	if proc, ok := r.lookup[pid.ID]; ok {
		return proc
	}
	return nil
}
