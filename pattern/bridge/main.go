package main

import (
	"pattern/bridge/brig"
	"net/http"
)

func main() {
	//建立一个客户端
	client := &brig.Client{http.DefaultClient}

	// cdn请求
	cdnReq := &brig.CdnRequest{}
	client.Query(cdnReq)

	// live请求
	liveReq := &brig.LiveRequest{}
	client.Query(liveReq)
}
