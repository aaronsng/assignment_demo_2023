// Code generated by Kitex v0.5.2. DO NOT EDIT.

package imservice

import (
	rpc "github.com/aaronsng/assignment_demo_2023/http-server/kitex_gen/rpc"
	server "github.com/cloudwego/kitex/server"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler rpc.IMService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
