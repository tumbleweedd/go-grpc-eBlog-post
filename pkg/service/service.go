package service

import (
	"context"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-post/pkg/pb"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-post/pkg/repository"
)

type Post interface {
	CreateNewPost(ctx context.Context, req *pb.CreateNewPostRequest) (*pb.CreateNewPostResponse, error)
	GetAllPosts(ctx context.Context, req *pb.GetAllPostsRequest) (*pb.GetAllPostsResponse, error)
	GetPostById(ctx context.Context, req *pb.GetPostByIdRequest) (*pb.GetPostByIdResponse, error)
	GetAllPostsByUserId(ctx context.Context, req *pb.GetAllPostsByUserIdRequest) (*pb.GetAllPostsByUserIdResponse, error)
	UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error)
	DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error)
}

type Service struct {
	Post
	pb.UnsafePostServiceServer
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Post: NewPostService(r.Post, r.Category, r.Tag),
	}
}
