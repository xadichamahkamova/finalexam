package service

import (
	"context"
	pb "user-service/genproto/userpb"
	"user-service/internal/storage"

	"github.com/google/uuid"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	Queries *storage.Queries
}

func NewUserService(repo *storage.Queries) *UserService {
	return &UserService{
		Queries: repo,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {

	user := storage.RegisterUserParams{
		UserName: req.UserName,
		Password: req.Password,
		Email:    req.Email,
	}
	data, err := s.Queries.RegisterUser(ctx, user)
	if err != nil {
		return nil, err
	}
	resp := pb.RegisterUserResponse{
		UserId:   data.ID.String(),
		UserName: data.UserName,
		Email:    data.Email,
	}
	return &resp, nil
}

func (s *UserService) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	user := storage.LoginUserParams{
		UserName: req.UserName,
		Password: req.Password,
	}
	data, err := s.Queries.LoginUser(ctx, user)
	if err != nil {
		return nil, err
	}
	resp := pb.LoginUserResponse{
		UserId: data.ID.String(),
		Email:  data.Email,
	}
	return &resp, nil
}

func (s *UserService) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {

	uuID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	data, err := s.Queries.GetUserById(ctx, uuID)
	if err != nil {
		return nil, err
	}
	resp := pb.GetUserByIdResponse{
		User: &pb.User{
			UserId:   data.ID.String(),
			UserName: data.UserName,
			Email:    data.Email,
		},
	}
	return &resp, nil
}

func (s *UserService) GetUsersList(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {

	resp := pb.GetUsersResponse{}
	users, err := s.Queries.GetUsersList(ctx)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		user := &pb.User{
			UserId:   user.ID.String(),
			UserName: user.UserName,
			Email:    user.Email,
		}
		resp.List = append(resp.List, user)
	}
	return &resp, nil
}
