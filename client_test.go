package gofcgi

import (
	"testing"
	"io/ioutil"
	"bytes"
	"time"
)

func TestClientGet(t *testing.T) {
	client := &Client{
		network: "tcp",
		address: "127.0.0.1:9000",
	}
	err := client.Connect()
	if err != nil {
		t.Fatal("connect err:", err.Error())
	}

	req := NewRequest()
	req.SetParams(map[string]string{
		"SCRIPT_FILENAME": "/Users/liuxiangchao/Documents/Projects/pp/apps/baleshop.ppk/index.php",
		"SERVER_SOFTWARE": "gofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgi/1.0.0",
		"REMOTE_ADDR":     "127.0.0.1",
		"QUERY_STRING":    "name=value&__ACTION__=/@wx",

		"SERVER_NAME":       "wx.balefm.cn",
		"SERVER_ADDR":       "127.0.0.1:80",
		"SERVER_PORT":       "80",
		"REQUEST_URI":       "/index.php?__ACTION__=/@wx",
		"DOCUMENT_ROOT":     "/Users/liuxiangchao/Documents/Projects/pp/apps/baleshop.ppk/",
		"GATEWAY_INTERFACE": "CGI/1.1",
		"REDIRECT_STATUS":   "200",
		"HTTP_HOST":         "wx.balefm.cn",

		"REQUEST_METHOD": "GET",
	})

	resp, err := client.Call(req)
	if err != nil {
		t.Fatal("do error:", err.Error())
	}

	t.Log("resp, status:", resp.StatusCode)
	t.Log("resp, status message:", resp.Status)
	t.Log("resp headers:", resp.Header)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("resp body:", string(data))
}

func TestClientGetAlive(t *testing.T) {
	client := &Client{
		network: "tcp",
		address: "127.0.0.1:9000",
	}
	client.KeepAlive()
	err := client.Connect()
	if err != nil {
		t.Fatal("connect err:", err.Error())
	}

	for i := 0; i < 10; i ++ {
		req := NewRequest()
		req.SetParams(map[string]string{
			"SCRIPT_FILENAME": "/Users/liuxiangchao/Documents/Projects/pp/apps/baleshop.ppk/index.php",
			"SERVER_SOFTWARE": "gofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgi/1.0.0",
			"REMOTE_ADDR":     "127.0.0.1",
			"QUERY_STRING":    "name=value&__ACTION__=/@wx",

			"SERVER_NAME":       "wx.balefm.cn",
			"SERVER_ADDR":       "127.0.0.1:80",
			"SERVER_PORT":       "80",
			"REQUEST_URI":       "/index.php?__ACTION__=/@wx",
			"DOCUMENT_ROOT":     "/Users/liuxiangchao/Documents/Projects/pp/apps/baleshop.ppk/",
			"GATEWAY_INTERFACE": "CGI/1.1",
			"REDIRECT_STATUS":   "200",
			"HTTP_HOST":         "wx.balefm.cn",

			"REQUEST_METHOD": "GET",
		})

		resp, err := client.Call(req)
		if err != nil {
			t.Fatal("do error:", err.Error())
		}

		t.Log("resp, status:", resp.StatusCode)
		t.Log("resp, status message:", resp.Status)
		t.Log("resp headers:", resp.Header)

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("resp body:", string(data))

		time.Sleep(1 * time.Second)
	}
}

func TestClientPost(t *testing.T) {
	client := &Client{
		network: "tcp",
		address: "127.0.0.1:9000",
	}
	err := client.Connect()
	if err != nil {
		t.Fatal("connect err:", err.Error())
	}

	req := NewRequest()
	req.SetParams(map[string]string{
		"SCRIPT_FILENAME": "/Users/liuxiangchao/Documents/Projects/pp/apps/baleshop.ppk/index.php",
		"SERVER_SOFTWARE": "gofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgigofcgi/1.0.0",
		"REMOTE_ADDR":     "127.0.0.1",
		"QUERY_STRING":    "name=value&__ACTION__=/@wx",

		"SERVER_NAME":       "wx.balefm.cn",
		"SERVER_ADDR":       "127.0.0.1:80",
		"SERVER_PORT":       "80",
		"REQUEST_URI":       "/index.php?__ACTION__=/@wx",
		"DOCUMENT_ROOT":     "/Users/liuxiangchao/Documents/Projects/pp/apps/baleshop.ppk/",
		"GATEWAY_INTERFACE": "CGI/1.1",
		"REDIRECT_STATUS":   "200",
		"HTTP_HOST":         "wx.balefm.cn",

		"REQUEST_METHOD": "POST",
		"CONTENT_TYPE":   "application/x-www-form-urlencoded",
	})

	r := bytes.NewReader([]byte("name12=value&hello=world&name13=value&hello4=world"))
	//req.SetParam("CONTENT_LENGTH", fmt.Sprintf("%d", r.Len()))
	req.SetBody(r, uint32(r.Len()))

	resp, err := client.Call(req)
	if err != nil {
		t.Fatal("do error:", err.Error())
	}

	t.Log("resp, status:", resp.StatusCode)
	t.Log("resp, status message:", resp.Status)
	t.Log("resp headers:", resp.Header)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("resp body:", string(data))
}
