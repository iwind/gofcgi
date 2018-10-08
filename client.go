package gofcgi

import (
	"errors"
	"net"
	"net/http"
	"sync"
	"time"
)

var ErrClientDisconnect = errors.New("[fcgi]lost connection to server")

type Client struct {
	isFree      bool
	isAvailable bool

	keepAlive bool

	network string
	address string
	conn    net.Conn

	locker sync.Mutex

	mock bool
}

func NewClient(network string, address string) *Client {
	return &Client{
		isFree:      true,
		isAvailable: false,
		network:     network,
		address:     address,
	}
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

	resp, err := req.CallOn(this.conn)

	// 如果失去连接，则重新连接
	if err != nil {
		if err == ErrClientDisconnect {
			this.Close()
		}
	}

	return resp, err
}

func (this *Client) Close() {
	this.isAvailable = false
	this.conn.Close()
	this.conn = nil
}

func (this *Client) Connect() error {
	// @TODO 设置并使用超时时间
	conn, err := net.Dial(this.network, this.address)
	if err != nil {
		this.isAvailable = false
		return err
	}

	this.conn = conn
	this.isAvailable = true

	return nil
}

func (this *Client) Mock() {
	this.mock = true
}
