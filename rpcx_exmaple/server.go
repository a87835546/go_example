package rpcx_exmaple

import "github.com/kataras/iris/v12"

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith int

func (t *Arith) Mul(ctx iris.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}
