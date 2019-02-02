package brig

import (
	"net/http"
)

/*
桥接模式
    适用于把抽象化与实现化解耦，使得二者可以独立变化。这种类型的设计模式属于结构型模式，
    它通过提供抽象化和实现化之间的桥接结构，来实现二者的解耦。

Bridge模式
基于类的最小设计原则，通过使用封装、聚合及继承等行为让不同的类承担不同的职责。
它的主要特点是把抽象(Abstraction)与行为实现(Implementation)分离开来，
从而可以保持各部分的独立性以及应对他们的功能扩展

*/

//请求接口 请求抽象
type Request interface {
	HttpReq() (*http.Request, error)
}

//客户端 基础
type Client struct {
	Client *http.Client
}

//请求查询
func (c *Client) Query(req Request) (resp *http.Response, err error) {
	//实现http请求
	httpRep, _ := req.HttpReq()
	//建立请求客户端连接
	resp, err = c.Client.Do(httpRep)
	return
}

//设定不同的请求 行为实现
//设置 cdn 请求
type CdnRequest struct {
}

//实现请求接口
func (cdn *CdnRequest) HttpReq() (*http.Request, error) {
	//发起一个请求
	return http.NewRequest("GET", "/cdn", nil)
}

//设置 live 请求
type LiveRequest struct {
}

//实现请求接口
func (live *LiveRequest) HttpReq() (*http.Request, error) {
	//发起一个请求
	return http.NewRequest("GET", "/live", nil)
}
