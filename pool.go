package gofcgi

import (
	"errors"
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

		// 第一个同步连接供使用，其余的可以异步建立连接
		if i == 0 {
			client.Connect()
		} else {
			go client.Connect()
		}
		pool.clients = append(pool.clients, client)
	}

	// 监控连接
	go func() {
		for {
			time.Sleep(10 * time.Second)
			for _, client := range pool.clients {
				if !client.isAvailable {
					client.Connect()
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
