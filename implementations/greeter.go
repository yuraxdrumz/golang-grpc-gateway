package implementations

import (
	"context"
	greeter "grpc_gateway/generated"
)

type Greeter struct {
	greeter.UnimplementedGreeterServer
}

func (g *Greeter) SayHello(ctx context.Context, in *greeter.HelloRequest) (*greeter.HelloReply, error) {
	return &greeter.HelloReply{Message: in.Message, Name: in.Name}, nil
}
