package sv

import "context"

/*
编写服务端
*/

//参数结构体
type Args struct {
	A int
	B int
}

//响应结构体
type Reply struct {
	C int
}

//定义了一个叫做 Arith 的 service， 并且为它实现了 Mul 方法。
//服务类型
type Arith struct {
}

// 服务方法
func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

