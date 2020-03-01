package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/dgrijalva/jwt-go"

	"github.com/RainJoe/ms/service/user/config"
	"github.com/RainJoe/ms/service/user/dao"

	pb "github.com/RainJoe/ms/service/user/proto/user"
)

//userService user micro service
type userService struct {
	dao *dao.Dao
}

//NewUserService create user service
func NewUserService() pb.UserServiceServer {
	return &userService{
		dao: dao.New(config.Cfg),
	}
}

//Create user rpc call
func (s *userService) Create(ctx context.Context, user *pb.User) (rsp *pb.Response, err error) {
	t := time.Now().UnixNano() / 1e6
	prettyID, err := uuid.NewUUID()
	if err != nil {
		return
	}
	user.PrettyId = prettyID.String()
	user.CreateTime = t
	user.UpdateTime = t
	if err = s.dao.CreateUser(ctx, user); err != nil {
		return
	}
	user.Password = ""
	rsp = &pb.Response{
		User: user,
	}
	return
}

//Get user rpc call
func (s *userService) Get(ctx context.Context, user *pb.User) (rsp *pb.Response, err error) {
	user, err = s.dao.GetUser(ctx, user)
	if err != nil {
		return
	}
	rsp = &pb.Response{
		User: user,
	}
	return
}

//GetAll user rpc call
func (s *userService) GetAll(ctx context.Context, req *pb.Request) (rsp *pb.Response, err error) {
	if req.PageNum == 0 || req.PageSize == 0 {
		req.PageNum = 1
		req.PageSize = 10
	}
	users, err := s.dao.GetAllUsers(ctx, req)
	if err != nil {
		return
	}
	rsp = &pb.Response{
		Users: users,
	}
	return
}

//Auth user with name and password rpc call
func (s *userService) Auth(ctx context.Context, user *pb.User) (ptoken *pb.Token, err error) {
	user, err = s.dao.ValidateUser(ctx, user.Name, user.Password)
	if err != nil {
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.PrettyId
	claims["name"] = user.Name
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(config.Cfg.Jwt.Key))
	if err != nil {
		return
	}
	if err = s.dao.SetUserToken2Redis(config.Cfg.Redis.TokenKey+user.Name, t); err != nil {
		return
	}
	ptoken = &pb.Token{
		Token: t,
	}
	return
}

//ValidateToken as its name says
func (s *userService) ValidateToken(ctx context.Context, token *pb.Token) (tokenOut *pb.Token, err error) {
	t, err := jwt.Parse(token.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.Jwt.Key), nil
	})
	if err != nil {
		return
	}
	if !t.Valid {
		err = errors.New("token invalid")
		return
	}
	claims := t.Claims.(jwt.MapClaims)
	if claims == nil {
		err = errors.New("token invalid")
	}
	var user pb.User
	user.PrettyId = claims["id"].(string)
	user.Name = claims["name"].(string)
	result, err := s.dao.GetUserTokenFromRedis(config.Cfg.Redis.TokenKey + user.Name)
	if err != nil {
		return
	}
	if result != t.Raw {
		err = errors.New("token invalid")
		return
	}
	tokenOut = &pb.Token{
		Valid: true,
	}
	return
}

//DeleteAuth delete token
func (s *userService) DeleteAuth(ctx context.Context, token *pb.Token) (rsp *pb.Response, err error) {
	t, err := jwt.Parse(token.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.Jwt.Key), nil
	})
	if err != nil {
		return
	}
	if !t.Valid {
		err = errors.New("token invalid")
		return
	}
	claims := t.Claims.(jwt.MapClaims)
	if claims == nil {
		err = errors.New("token invalid")
		return
	}
	var user pb.User
	user.PrettyId = claims["id"].(string)
	user.Name = claims["name"].(string)
	if err = s.dao.DelUserTokenFromRedis(config.Cfg.Redis.TokenKey + user.Name); err != nil {
		return
	}
	return
}
