package objectpool

import (
	"fmt"
	"sync"
)

type pool struct {
	idle     []iPoolObject
	active   []iPoolObject
	capacity int
	muLock   *sync.Mutex
}

func initPool(poolObject []iPoolObject) (*pool, error) {
	if len(poolObject) == 0 {
		return nil, fmt.Errorf("cannot create a pool of o length")
	}
	active := make([]iPoolObject, 0)
	pool := &pool{
		idle:     poolObject,
		active:   active,
		capacity: len(poolObject),
		muLock:   new(sync.Mutex),
	}
	return pool, nil
}

func (p *pool) loan() (iPoolObject, error) {
	p.muLock.Lock()
	defer p.muLock.Unlock()
	if len(p.idle) == 0 {
		return nil, fmt.Errorf("no idle object free,Please request after sometime")
	}
	obj := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, obj)
	fmt.Printf("Loan Pool Object With Id :%s\n", obj.getID())
	return obj, nil
}

func (p *pool) receive(target iPoolObject) error {
	p.muLock.Lock()
	defer p.muLock.Unlock()
	err := p.remove(target)
	if err != nil {
		return err
	}
	p.idle = append(p.idle, target)
	fmt.Printf("Return Pool Object with ID: %s\n", target.getID())
	return nil
}

func (p *pool) remove(target iPoolObject) error {
	currentActiveLength := len(p.active)
	for i, obj := range p.active {
		if obj.getID() == target.getID() {
			p.active[currentActiveLength-1], p.active[i] = p.active[i], p.active[currentActiveLength-1]
			p.active = p.active[:currentActiveLength-1]
			return nil
		}
	}

	return fmt.Errorf("targe pool object doesn't belong to the pool")
}
