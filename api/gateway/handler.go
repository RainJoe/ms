package gateway

import (
	"github.com/RainJoe/ms/api/config"
	pb "github.com/RainJoe/ms/service/user/proto/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommonResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Version string `json:"version"`
	Data interface{} `json:"data,omitempty"`
}


//Handler provides services
func Handler(userService pb.UserServiceClient) *gin.Engine {
	router := gin.Default()

	router.POST("/v1/users", createUser(userService))
	router.GET("/v1/users", getAllUsers(userService))
	router.GET("/v1/user", getUser(userService))
	router.POST("/v1/login", login(userService))

	return router
}

func RespSuccess(c *gin.Context, data interface{}) {
	Resp(c, 0, "Success", data)
}

func RespError(c *gin.Context, errCode ErrCode, data interface{}) {
	Resp(c, int(errCode), ErrMessages[errCode], data)
}
func Resp(c *gin.Context, code int, message string, data interface{}) {
	cm := CommonResponse{}
	cm.Code = code
	cm.Message = message
	cm.Version = config.Cfg.Version
	cm.Data = data
	c.JSON(http.StatusOK, cm)
}