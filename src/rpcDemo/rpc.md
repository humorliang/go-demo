1. golang服务端的方法名注册和调用保持一致，
格式为： 结构体名.方法名

2. 能够注册的方法必须满足指定的函数签名：
　　func (t *T) MethodName(argType T1, replyType *T2) error 。
入参有且只有两个，第一个是被调用函数的输入参数，第二个是被调用函数的输出参数。
返回值只有一个error类型。这三个参数一个都不能变。


// server
1. 定义结构体和服务方法
2. 注册服务
3. 监听端口
4. 建立链接 event loop: accept()
// 连接处理
5. rpc.ServerConn(con)
