package main

import (
	"net/http"
	"io"
	"log"
)

/*
建立一个https连接：Create a Basic HTTPS Server (using TLS)

Step 1: Commands to make self-signed certificates
```
openssl genrsa -out https-server.key 2048
openssl ecparam -genkey -name secp384r1 -out https-server.key
```
Finally self-sign the certificate and specify information like country, name and email:
```
openssl req -new -x509 -sha256 -key https-server.key -out https-server.crt -days 3650
```
*/
func main() {
	//响应函数
	http.HandleFunc("/", ExampleHandler)
	log.Println("***** TLS server Listen On Port:8080 *****")
	//启动监听
	err := http.ListenAndServeTLS(":8080", "https-server.crt", "https-server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok"}`)
}
