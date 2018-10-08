# description
a **golang client for fastcgi**, support connection pool and easy to use.

# pool usage
~~~golang
// retrieve shared pool
pool := gofcgi.SharedPool("tcp", "127.0.0.1", 16)
client, err := pool.Client()
if err != nil {
    return
}

// create a request
req := gofcgi.NewRequest()
params = map[string]string{
	"SCRIPT_FILENAME": "[PATH TO YOUR SCRIPT]/index.php",
	"SERVER_SOFTWARE": "gofcgi/1.0.0",
	"REMOTE_ADDR":     "127.0.0.1",
	"QUERY_STRING":    "NAME=VALUE",

	"SERVER_NAME":       "example.com",
	"SERVER_ADDR":       "127.0.0.1:80",
	"SERVER_PORT":       "80",
	"REQUEST_URI":       "/index.php",
	"DOCUMENT_ROOT":     "[PATH TO YOUR SCRIPT]",
	"GATEWAY_INTERFACE": "CGI/1.1",
	"REDIRECT_STATUS":   "200",
	"HTTP_HOST":         "example.com",

	"REQUEST_METHOD": "POST",                              // for post method
	"CONTENT_TYPE":   "application/x-www-form-urlencoded", // for post
}

req.SetTimeout(5 * time.Second)
req.SetParams(params)

// set request body
r := bytes.NewReader([]byte("name=lu&age=20"))
req.SetBody(r, uint32(r.Len()))

// call request
resp, err := client.Call(req)
if err != nil {
    return
}

// read data from response
data, err := ioutil.ReadAll(resp.Body)
if err != nil {
    return
}
log.Println("resp body:", string(data))
~~~