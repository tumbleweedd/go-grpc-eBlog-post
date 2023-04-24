package client

import (
	"context"
	"fmt"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-post/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CommentServiceClient struct {
	Client pb.CommentServiceClient
}

func InitCommentServiceClient(url string) CommentServiceClient {
	cc, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := CommentServiceClient{
		Client: pb.NewCommentServiceClient(cc),
	}

	return c
}

func (c *CommentServiceClient) GetCommentByPostId(postId int) (*pb.GetCommentsByPostIdResponse, error) {
	req := &pb.GetCommentsByPostIdRequest{PostId: int64(postId)}

	return c.Client.GetCommentsByPostId(context.Background(), req)
}
