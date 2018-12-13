package main

import (
	"net/url"
	"fmt"
	"strings"
)

/*
url解析：
URL 提供了一个统一资源定位方式。
它包含了一个 scheme，认证信息，主机名，端口，路径，查询参数和片段。
*/
func main() {
	s := "postgres://user:123456@host.com:1234/path?k=v#f"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("url:", u)                          //url: postgres://user:123456@host.com:1234/path?k=v#f
	fmt.Println("scheme: ", u.Scheme)               //postgres
	fmt.Println("所有认证信息：", u.User)                  //所有认证信息： user:123456
	fmt.Println("认证信息Username：", u.User.Username()) //认证信息Username： user

	p, _ := u.User.Password()
	fmt.Println(p) //123456

	//host信息
	fmt.Println(u.Host) //host.com:1234
	host := strings.Split(u.Host, ":")
	fmt.Println(host[0])
	fmt.Println(host[1]) //1234

	//path
	fmt.Println(u.Path) //  /path
	//url片段
	fmt.Println(u.Fragment) //f
	//查询query参数
	fmt.Println(u.RawQuery) // k=v
	//将查询参数转化为一个map
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)         //map[k:[v]]
	fmt.Println(m["k"][0]) //v

}
