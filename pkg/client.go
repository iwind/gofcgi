package pkg

import (
	"errors"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

var ErrClientDisconnect = errors.New("lost connection to server")

type Client struct {
	isFree      bool
	isAvailable bool

	keepAlive bool

	network string
	address string
	conn    net.Conn

	locker sync.Mutex

	expireTime   time.Time
	expireLocker sync.Mutex

	mock bool
}

func NewClient(network string, address string) *Client {
	client := &Client{
		isFree:      true,
		isAvailable: false,
		network:     network,
		address:     address,
		expireTime:  time.Now().Add(86400 * time.Second),
	}

	// deal with expireTime
	go func() {
		for {
			time.Sleep(1 * time.Second)
			if time.Since(client.expireTime) > 0 {
				_ = client.conn.Close()

				client.expireLocker.Lock()
				client.expireTime = time.Now().Add(86400 * time.Second)
				client.expireLocker.Unlock()
			}
		}
	}()
	return client
}

func (this *Client) KeepAlive() {
	this.keepAlive = true
}

func (this *Client) Call(req *Request) (resp *http.Response, stderr []byte, err error) {
	this.isFree = false

	this.locker.Lock()

	if this.keepAlive && this.conn == nil {
		err := this.Connect()
		if err != nil {
			this.locker.Unlock()
			return nil, nil, err
		}
	}

	if this.keepAlive {
		req.keepAlive = true
	}

	defer func() {
		if this.mock {
			time.Sleep(1 * time.Second)
		}
		this.isFree = true
		this.locker.Unlock()
	}()

	if this.conn == nil {
		return nil, nil, errors.New("no connection to server")
	}

	if req.timeout > 0 {
		this.beforeTime(req.timeout)
	}
	resp, stderr, err = req.CallOn(this.conn)
	this.endTime()

	// if lost connection, retry
	if err != nil {
		log.Println("[gofcgi]" + err.Error())

		if err == ErrClientDisconnect {
			// retry again
			this.Close()
			err = this.Connect()
			if err == nil {
				if req.timeout > 0 {
					this.beforeTime(req.timeout)
				}
				resp, stderr, err = req.CallOn(this.conn)
				this.endTime()
			} else {
				log.Println("[gofcgi]again:" + err.Error())
				this.Close()
			}
		}
	}

	return resp, stderr, err
}

func (this *Client) Close() {
	this.isAvailable = false
	if this.conn != nil {
		_ = this.conn.Close()
	}
	this.conn = nil
}

func (this *Client) Connect() error {
	this.isAvailable = false

	// @TODO set timeout
	conn, err := net.Dial(this.network, this.address)
	if err != nil {
		log.Println("[gofcgi]" + err.Error())
		return err
	}

	this.conn = conn
	this.isAvailable = true

	return nil
}

func (this *Client) Mock() {
	this.mock = true
}

func (this *Client) beforeTime(timeout time.Duration) {
	this.expireLocker.Lock()
	this.expireTime = time.Now().Add(timeout)
	this.expireLocker.Unlock()
}

func (this *Client) endTime() {
	this.expireLocker.Lock()
	this.expireTime = time.Now().Add(86400 * time.Second)
	this.expireLocker.Unlock()
}
