package demo1

import (
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/kataras/iris/v12"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"test/grpc_service/app2"
)

type GRPCClient struct {
}

type server struct {
	app2.UnimplementedTestServer
}

func (s *server) AddProduct(ctx iris.Context, req *app2.Req) (res *app2.Res, err error) {
	res = &app2.Res{}
	_, err = uuid.NewV4()
	if err != nil {
		return res, status.Errorf(codes.Internal, "err while generate the uuid ", err)
	}

	return nil, err
}
