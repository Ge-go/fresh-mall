package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"mall-api/user-web/global/response"
	"mall-api/user-web/proto"
	"net/http"
	"time"
)

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	// 将grpc的code转换成http的状态码
	if e, ok := status.FromError(err); ok {
		switch e.Code() {
		case codes.NotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"msg": e.Message(),
			})
		case codes.InvalidArgument:
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "invalidate param",
			})
		case codes.AlreadyExists:
			c.JSON(http.StatusAlreadyReported, gin.H{
				"msg": "already exists",
			})
		case codes.Internal:
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "internal server error",
			})
		case codes.Unavailable:
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "User service not available",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "unknown error",
			})
		}
	}
}

func GetUserList(ctx *gin.Context) {
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[grpc.Dial] conn err", "msg", err.Error())
	}

	client := proto.NewUserClient(conn)
	userList, err := client.GetUserList(ctx, &proto.PageInfo{
		Pn:    1,
		PSize: 3,
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] get user list err")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	resp := make([]response.UserResponse, 0, len(userList.Data))
	for _, v := range userList.Data {
		resp = append(resp, response.UserResponse{
			Id:       v.Id,
			NickName: v.NickName,
			Birthday: time.Unix(int64(v.BirthDay), 0).Format("2006-01-02"),
			Mobile:   v.Mobile,
			Gender:   v.Gender,
		})
	}

	ctx.JSON(http.StatusOK, resp)
}
