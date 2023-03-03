package service

import (
	"context"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-post/pkg/model"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-post/pkg/pb"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-post/pkg/repository"
	"net/http"
)

type PostService struct {
	postRepo     repository.Post
	categoryRepo repository.Category
	tagRepo      repository.Tag
}

func NewPostService(postRepo repository.Post, categoryRepo repository.Category, tagRepo repository.Tag) *PostService {
	return &PostService{
		postRepo:     postRepo,
		categoryRepo: categoryRepo,
		tagRepo:      tagRepo,
	}
}

func (postService *PostService) CreateNewPost(ctx context.Context, req *pb.CreateNewPostRequest) (*pb.CreateNewPostResponse, error) {
	categoryId, err := postService.categoryRepo.GetCategoryIdByName(req.Data.GetCategory())
	if err != nil {
		return &pb.CreateNewPostResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	post := model.PostDTO{
		Body:     req.Data.GetBody(),
		Head:     req.Data.GetHead(),
		Category: req.Data.GetCategory(),
		Tags:     req.Data.GetTags(),
	}

	err = postService.postRepo.CreateNewPost(categoryId, int(req.UserId), post)
	if err != nil {
		return &pb.CreateNewPostResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	return &pb.CreateNewPostResponse{
		Status: http.StatusOK,
		Head:   post.Head,
		Body:   post.Body,
	}, nil
}

func (postService *PostService) GetAllPosts(ctx context.Context, req *pb.GetAllPostsRequest) (*pb.GetAllPostsResponse, error) {
	posts, err := postService.postRepo.GetAllPosts()
	if err != nil {
		return &pb.GetAllPostsResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	data := make([]*pb.PostData, 0, len(posts))
	postsDTO, err := generatePostDTO(posts, postService)
	if err != nil {
		return &pb.GetAllPostsResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	for _, post := range postsDTO {
		data = append(data, &pb.PostData{
			Body:     post.Body,
			Head:     post.Head,
			Category: post.Category,
			Tags:     post.Tags,
		})
	}

	return &pb.GetAllPostsResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (postService *PostService) GetPostById(ctx context.Context, req *pb.GetPostByIdRequest) (*pb.GetPostByIdResponse, error) {
	post, err := postService.postRepo.GetPostById(int(req.GetPostId()))
	if err != nil {
		return &pb.GetPostByIdResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	postCategory, err := postService.categoryRepo.FindCategoryById(post.CategoryId)
	postTags, err := postService.tagRepo.GetPostTagsByPostId(post.Id)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	data := &pb.PostData{
		Head:     post.Head,
		Body:     post.Body,
		Category: postCategory.Name,
		Tags:     getTagsInSlice(postTags),
	}

	return &pb.GetPostByIdResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (postService *PostService) GetAllPostsByUserId(ctx context.Context, req *pb.GetAllPostsByUserIdRequest) (*pb.GetAllPostsByUserIdResponse, error) {
	userPosts, err := postService.postRepo.GetPostsByUserId(int(req.GetUserId()))
	if err != nil {
		return &pb.GetAllPostsByUserIdResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	data := make([]*pb.PostData, 0, len(userPosts))
	userPostsDTO, err := generatePostDTO(userPosts, postService)
	if err != nil {
		return &pb.GetAllPostsByUserIdResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	for _, post := range userPostsDTO {
		data = append(data, &pb.PostData{
			Body:     post.Body,
			Head:     post.Head,
			Category: post.Category,
			Tags:     post.Tags,
		})
	}

	return &pb.GetAllPostsByUserIdResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (postService *PostService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (postService *PostService) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	//TODO implement me
	panic("implement me")
}

func generatePostDTO(posts []model.Post, postService *PostService) ([]model.PostDTO, error) {
	var postsDTO = make([]model.PostDTO, len(posts), len(posts))
	var tags []model.Tag

	for index, post := range posts {
		postsDTO[index].Head = post.Head
		postsDTO[index].Body = post.Body

		category, err := postService.categoryRepo.FindCategoryById(post.CategoryId)
		if err != nil {
			return nil, err
		}
		postsDTO[index].Category = category.Name

		tags, err = postService.tagRepo.GetPostTagsByPostId(post.Id)
		if err != nil {
			return nil, err
		}
		postsDTO[index].Tags = getTagsInSlice(tags)
	}

	return postsDTO, nil
}

func getTagsInSlice(tags []model.Tag) []string {
	var tagInString []string
	for _, tag := range tags {
		tagInString = append(tagInString, tag.Name)
	}
	return tagInString
}
