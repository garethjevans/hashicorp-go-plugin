package greeter

import (
	"github.com/hashicorp/go-plugin"
	"net/rpc"
)

// This is the implementation of plugin.Plugin so we can serve/consume this
//
// This has two methods: Server must return an RPC server for this plugin
// type. We construct a GreeterRPCServer for this.
//
// Client must return an implementation of our interface that communicates
// over an RPC client. We return GreeterRPC for this.
//
// Ignore MuxBroker. That is used to create more multiplexed streams on our
// plugin connection and is a more advanced use case.
type GreeterPlugin struct {
	// Impl Injection
	Impl Greeter
}

func (p *GreeterPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &GreeterRPCServer{Impl: p.Impl}, nil
}

func (GreeterPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &GreeterRPC{client: c}, nil
}
