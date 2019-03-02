package main

import (
	"flag" //flag 包实现了命令行参数的解析。
	"html/template"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

//定义命令行参数  名字  默认值   提示
var addr = flag.String("addr", "localhost:8080", "http service address")

// 定义ws协议的选项
var upgrader = websocket.Upgrader{
	//ws跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echo(w http.ResponseWriter, r *http.Request) {
	//创建ws连接
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	//延迟函数关闭
	defer c.Close()
	//循环监听
	for {
		//读取消息  返回：消息类型  消息字节类型slice  错误
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		//消息循环监听日志
		log.Printf("recv: %s", message)
		//写入消息 字节类型的slice
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

//返回主页路由函数
func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func main() {
	//解析注册命令行参数
	flag.Parse()
	log.SetFlags(0)
	//ws消息接受路由
	http.HandleFunc("/echo", echo)
	//ws界面访问路由
	http.HandleFunc("/", home)
	//监听地址
	log.Fatal(http.ListenAndServe(*addr, nil))
}

//返回一个模板信息
var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            console.log("OPEN");
        }
        ws.onclose = function(evt) {
            console.log("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            console.log("RESPONSE: " + evt.car);
        }
        ws.onerror = function(evt) {
            console.log("ERROR: " + evt.car);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        console.log("SEND: " + input.value);
        ws.send(input.value);
		input.value=""
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
