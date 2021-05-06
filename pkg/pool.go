package pkg

import (
	"errors"
	"log"
	"sync"
	"time"
)

var pools = map[string]*Pool{} // fullAddress => *Pool
var poolsLocker = sync.Mutex{}

type Pool struct {
	size    uint
	timeout time.Duration
	clients []*Client
	locker  sync.Mutex
}

func SharedPool(network string, address string, size uint) *Pool {
	poolsLocker.Lock()
	defer poolsLocker.Unlock()

	fullAddress := network + "//" + address
	pool, found := pools[fullAddress]
	if found {
		return pool
	}

	if size == 0 {
		size = 8
	}

	pool = &Pool{
		size: size,
	}

	for i := uint(0); i < size; i++ {
		client := NewClient(network, address)
		client.KeepAlive()

		// prepare one for first request, and left for async request
		if i == 0 {
			err := client.Connect()
			if err != nil {
				log.Println("[gofcgi]" + err.Error())
			}
		} else {
			go func() {
				err := client.Connect()
				if err != nil {
					log.Println("[gofcgi]" + err.Error())
				}
			}()
		}
		pool.clients = append(pool.clients, client)
	}

	// watch connections
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		for range ticker.C {
			for _, client := range pool.clients {
				if !client.isAvailable {
					_ = client.Connect()
				}
			}
		}
	}()

	pools[fullAddress] = pool

	return pool
}

func (this *Pool) Client() (*Client, error) {
	this.locker.Lock()
	defer this.locker.Unlock()

	if len(this.clients) == 0 {
		return nil, errors.New("no available clients to use")
	}

	// find a free one
	for _, client := range this.clients {
		if client.isAvailable && client.isFree {
			return client, nil
		}
	}

	// find available on
	for _, client := range this.clients {
		if client.isAvailable {
			return client, nil
		}
	}

	// use first one
	err := this.clients[0].Connect()
	if err == nil {
		return this.clients[0], nil
	}

	return nil, errors.New("no available clients to use")
}
