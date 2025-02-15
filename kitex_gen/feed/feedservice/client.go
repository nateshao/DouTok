// Code generated by Kitex v0.4.4. DO NOT EDIT.

package feedservice

import (
	"context"
	feed "github.com/TremblingV5/DouTok/kitex_gen/feed"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetUserFeed(ctx context.Context, Req *feed.DouyinFeedRequest, callOptions ...callopt.Option) (r *feed.DouyinFeedResponse, err error)
	GetVideoById(ctx context.Context, Req *feed.VideoIdRequest, callOptions ...callopt.Option) (r *feed.Video, err error)
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
	return &kFeedServiceClient{
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

type kFeedServiceClient struct {
	*kClient
}

func (p *kFeedServiceClient) GetUserFeed(ctx context.Context, Req *feed.DouyinFeedRequest, callOptions ...callopt.Option) (r *feed.DouyinFeedResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserFeed(ctx, Req)
}

func (p *kFeedServiceClient) GetVideoById(ctx context.Context, Req *feed.VideoIdRequest, callOptions ...callopt.Option) (r *feed.Video, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVideoById(ctx, Req)
}
