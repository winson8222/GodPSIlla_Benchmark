// Code generated by Kitex v0.5.2. DO NOT EDIT.

package posterservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	postersapi "testposterservice/kitex_gen/postersapi"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Getuniqueusernames8020(ctx context.Context, req *postersapi.Request, callOptions ...callopt.Option) (r *postersapi.Response, err error)
	Getuniqueusernames2080(ctx context.Context, req *postersapi.Request, callOptions ...callopt.Option) (r *postersapi.Response, err error)
	Getuniqueusernames8099920(ctx context.Context, req *postersapi.Request, callOptions ...callopt.Option) (r *postersapi.Response, err error)
	Getuniqueusernames2099980(ctx context.Context, req *postersapi.Request, callOptions ...callopt.Option) (r *postersapi.Response, err error)
	Getuniqueusernames80k20k(ctx context.Context, req *postersapi.Request, callOptions ...callopt.Option) (r *postersapi.Response, err error)
	Getuniqueusernames20k80k(ctx context.Context, req *postersapi.Request, callOptions ...callopt.Option) (r *postersapi.Response, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kPosterServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kPosterServiceClient struct {
	*kClient
}

func (p *kPosterServiceClient) Getuniqueusernames8020(ctx context.Context, req *postersapi.Request, callOptions ...callopt.Option) (r *postersapi.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Getuniqueusernames8020(ctx, req)
}

func (p *kPosterServiceClient) Getuniqueusernames2080(ctx context.Context, req *postersapi.Request, callOptions ...callopt.Option) (r *postersapi.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Getuniqueusernames2080(ctx, req)
}

func (p *kPosterServiceClient) Getuniqueusernames8099920(ctx context.Context, req *postersapi.Request, callOptions ...callopt.Option) (r *postersapi.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Getuniqueusernames8099920(ctx, req)
}

func (p *kPosterServiceClient) Getuniqueusernames2099980(ctx context.Context, req *postersapi.Request, callOptions ...callopt.Option) (r *postersapi.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Getuniqueusernames2099980(ctx, req)
}

func (p *kPosterServiceClient) Getuniqueusernames80k20k(ctx context.Context, req *postersapi.Request, callOptions ...callopt.Option) (r *postersapi.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Getuniqueusernames80k20k(ctx, req)
}

func (p *kPosterServiceClient) Getuniqueusernames20k80k(ctx context.Context, req *postersapi.Request, callOptions ...callopt.Option) (r *postersapi.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Getuniqueusernames20k80k(ctx, req)
}
