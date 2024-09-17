package service

import (
	"context"
	"errors"
	pb "user-service/genproto/userpb"
	hash "user-service/internal/hash"
	"user-service/internal/storage"
	"user-service/logger"

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

	logger.Info("RegisterUser: Registering new user - UserName: ", req.UserName)

	user := storage.RegisterUserParams{
		UserName: req.UserName,
		Password: req.Password,
		Email:    req.Email,
	}

	passwordHash, err := hash.HashPassword(user.Password)
	if err != nil {
		logger.Error("RegisterUser: Error while hashing password - ", err)
		return nil, err
	}

	user.Password = passwordHash
	data, err := s.Queries.RegisterUser(ctx, user)
	if err != nil {
		logger.Error("RegisterUser: Error registering user - ", err)
		return nil, err
	}

	resp := pb.RegisterUserResponse{
		UserId:   data.ID.String(),
		UserName: data.UserName,
		Email:    data.Email,
	}

	logger.Info("RegisterUser: Successfully registered user - UserId: ", resp.UserId)
	return &resp, nil
}

func (s *UserService) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	logger.Info("LoginUser: Attempting login - UserName: ", req.UserName)

	user, err := s.Queries.GetUserByUserName(ctx, req.UserName)
	if err != nil {
		logger.Error("LoginUser: Error fetching user - ", err)
		return nil, err
	}

	isValid := hash.CheckPasswordHash(req.Password, user.Password)
	if !isValid {
		logger.Error("LoginUser: Invalid password")
		return nil, errors.New("invalid password")
	}

	resp := pb.LoginUserResponse{
		UserId: user.ID.String(),
		Email:  user.Email,
	}

	logger.Info("LoginUser: Successfully logged in user - UserId: ", resp.UserId)
	return &resp, nil
}

func (s *UserService) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {

	logger.Info("GetUserById: Fetching user by ID - UserId: ", req.UserId)

	uuID, err := uuid.Parse(req.UserId)
	if err != nil {
		logger.Error("GetUserById: Invalid UUID - ", err)
		return nil, err
	}

	data, err := s.Queries.GetUserById(ctx, uuID)
	if err != nil {
		logger.Error("GetUserById: Error retrieving user - ", err)
		return nil, err
	}

	resp := pb.GetUserByIdResponse{
		User: &pb.User{
			UserId:   data.ID.String(),
			UserName: data.UserName,
			Email:    data.Email,
		},
	}

	logger.Info("GetUserById: Successfully retrieved user - UserId: ", resp.User.UserId)
	return &resp, nil
}

func (s *UserService) GetUsersList(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {

	logger.Info("GetUsersList: Fetching users list")

	resp := pb.GetUsersResponse{}
	users, err := s.Queries.GetUsersList(ctx)
	if err != nil {
		logger.Error("GetUsersList: Error retrieving users list - ", err)
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

	logger.Info("GetUsersList: Successfully retrieved users list - Count: ", len(resp.List))
	return &resp, nil
}
