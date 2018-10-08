package gofcgi

import (
	"testing"
	"time"
	"errors"
	"io/ioutil"
)

func TestSharedPool(t *testing.T) {
	pool := SharedPool("tcp", "127.0.0.1:9000", 8)

	time.Sleep(100 * time.Millisecond)

	for i := 0; i < 3; i ++ {
		go func() {
			client, err := pool.Client()
			if err != nil {
				t.Fatal(errors.New("client should not be nil"))
			}
			t.Logf("client:%p", client)

			req := NewRequest()
			req.KeepAlive()

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
				t.Log(err.Error())
			} else {
				bodyBytes, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Log(err.Error())
				} else {
					t.Log(string(bodyBytes))
				}
			}
		}()
	}

	time.Sleep(2 * time.Second)
}
