// Code generated by Kitex v1.4.1. DO NOT EDIT.

package echo

import (
	"context"

	"github.com/cloudwego/hertz-examples/tracer/kitex/kitex_gen/api"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Echo(ctx context.Context, req *api.Request, callOptions ...callopt.Option) (r *api.Response, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(psm string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(psm))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kEchoClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(psm string, opts ...client.Option) Client {
	kc, err := NewClient(psm, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kEchoClient struct {
	*kClient
}

func (p *kEchoClient) Echo(ctx context.Context, req *api.Request, callOptions ...callopt.Option) (r *api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Echo(ctx, req)
}
