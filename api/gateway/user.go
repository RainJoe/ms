package gateway

import (
	"context"
	pb "github.com/RainJoe/ms/service/user/proto/user"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type UserGateway interface {
	CreateUser(c *gin.Context)
}

func createUser(userService pb.UserServiceClient) func(c *gin.Context) {
	return func (c *gin.Context) {
		req := struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{}
		if err := c.BindJSON(&req); err != nil {
			log.Printf("bind params failed: %v", err)
			RespError(c, ErrValidateParams, nil)
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := userService.Create(ctx, &pb.User{Name: req.Username, Password: req.Password})
		if err != nil {
			log.Printf("call Create failed: %v", err)
			RespError(c, ErrFailedToCallUserService, nil)
			return
		}
		RespSuccess(c, r)
	}
}

func getAllUsers(userService pb.UserServiceClient) func (c *gin.Context) {
	return func(c *gin.Context) {
		req := struct {
			PageNum  int64 `json:"page_num,omitempty"`
			PageSize int64 `json:"page_size,omitempty"`
		}{}
		if err := c.BindJSON(&req); err != nil {
			log.Printf("bind params failed: %v", err)
			RespError(c, ErrValidateParams, nil)
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := userService.GetAll(ctx, &pb.Request{PageNum: req.PageNum, PageSize: req.PageSize})
		if err != nil {
			log.Printf("call GetAll failed: %v", err)
			RespError(c, ErrFailedToCallUserService, nil)
			return
		}
		RespSuccess(c, r)
	}
}

func getUser(userService pb.UserServiceClient) func (c *gin.Context) {
	return func(c *gin.Context) {
		req := struct {
			UID string `json:"u_id"`
		}{}
		if err := c.BindJSON(&req); err != nil {
			log.Printf("bind params failed: %v", err)
			RespError(c, ErrValidateParams, nil)
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := userService.Get(ctx, &pb.User{PrettyId: req.UID})
		if err != nil {
			log.Printf("call GetAll failed: %v", err)
			RespError(c, ErrFailedToCallUserService, nil)
			return
		}
		RespSuccess(c, r)
	}
}

func login(userService pb.UserServiceClient) func (c *gin.Context) {
	return func(c *gin.Context) {
		req := struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{}
		if err := c.BindJSON(&req); err != nil {
			log.Printf("bind params failed: %v", err)
			RespError(c, ErrValidateParams, nil)
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := userService.Auth(ctx, &pb.User{Name: req.Username, Password: req.Password})
		if err != nil {
			log.Printf("call Auth failed: %v", err)
			RespError(c, ErrFailedToCallUserService, nil)
			return
		}
		RespSuccess(c, r)
	}
}