package gofcgi

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

	// 处理超时
	go func() {
		for {
			time.Sleep(1 * time.Second)
			if time.Since(client.expireTime) > 0 {
				client.conn.Close()

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

func (this *Client) Call(req *Request) (*http.Response, error) {
	this.isFree = false

	this.locker.Lock()

	if this.keepAlive && this.conn == nil {
		err := this.Connect()
		if err != nil {
			this.locker.Unlock()
			return nil, err
		}
	}

	if this.keepAlive {
		req.keepAlive = true
	}

	defer func() {
		if this.mock {
			time.Sleep(1 * time.Second) // 模拟占用连接
		}
		this.isFree = true
		this.locker.Unlock()
	}()

	if this.conn == nil {
		return nil, errors.New("no connection to server")
	}

	if req.timeout > 0 {
		this.beforeTime(req.timeout)
	}
	resp, err := req.CallOn(this.conn)
	this.endTime()

	// 如果失去连接，则重新连接
	if err != nil {
		log.Println("[gofcgi]" + err.Error())

		if err == ErrClientDisconnect {
			// 重试一次
			this.Close()
			err = this.Connect()
			if err == nil {
				if req.timeout > 0 {
					this.beforeTime(req.timeout)
				}
				resp, err = req.CallOn(this.conn)
				this.endTime()
			} else {
				log.Println("[gofcgi]again:" + err.Error())
				this.Close()
			}
		}
	}

	return resp, err
}

func (this *Client) Close() {
	this.isAvailable = false
	if this.conn != nil {
		this.conn.Close()
	}
	this.conn = nil
}

func (this *Client) Connect() error {
	this.isAvailable = false

	// @TODO 设置并使用超时时间
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
