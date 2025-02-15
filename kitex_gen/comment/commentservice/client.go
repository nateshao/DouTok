// Code generated by Kitex v0.3.4. DO NOT EDIT.

package commentservice

import (
	"context"
	"github.com/TremblingV5/DouTok/kitex_gen/comment"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CommentAction(ctx context.Context, Req *comment.DouyinCommentActionRequest, callOptions ...callopt.Option) (r *comment.DouyinCommentActionResponse, err error)
	CommentList(ctx context.Context, Req *comment.DouyinCommentListRequest, callOptions ...callopt.Option) (r *comment.DouyinCommentListResponse, err error)
}

// NewClient creates a rpc for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kCommentServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a rpc for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCommentServiceClient struct {
	*kClient
}

func (p *kCommentServiceClient) CommentAction(ctx context.Context, Req *comment.DouyinCommentActionRequest, callOptions ...callopt.Option) (r *comment.DouyinCommentActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentAction(ctx, Req)
}

func (p *kCommentServiceClient) CommentList(ctx context.Context, Req *comment.DouyinCommentListRequest, callOptions ...callopt.Option) (r *comment.DouyinCommentListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentList(ctx, Req)
}
