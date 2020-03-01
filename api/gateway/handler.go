package gateway

import (
	pb "github.com/RainJoe/ms/service/user/proto/user"
	"github.com/gin-gonic/gin"
)

//Handler provides services
func Handler(userService pb.UserServiceClient) *gin.Engine {
	router := gin.Default()

	router.POST("/v1/users", createUser(userService))
	router.GET("/v1/users", getAllUsers(userService))
	router.GET("/v1/user", getUser(userService))
	router.POST("/v1/login", login(userService))

	return router
}
