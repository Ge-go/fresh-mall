package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	jwtmodel "mall-api/user-web/global/jwt"
	"mall-api/user-web/middleware"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"mall-api/user-web/forms"
	"mall-api/user-web/global"
	"mall-api/user-web/global/response"
	"mall-api/user-web/proto"
	"mall-api/user-web/utils"
)

// GetUserList 获取用户列表
func GetUserList(ctx *gin.Context) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[grpc.Dial] conn err", "msg", err.Error())
	}

	// todo validate
	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)

	client := proto.NewUserClient(conn)
	userList, err := client.GetUserList(ctx, &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] get user list err")
		utils.HandleGrpcErrorToHttp(err, ctx)
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

// PassWordLogin 密码登录(尚未注册)
func PassWordLogin(ctx *gin.Context) {
	//表单验证
	passwordLoginForm := forms.PassWordLoginForm{}
	if err := ctx.ShouldBind(&passwordLoginForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[grpc.Dial] conn err", "msg", err.Error())
	}

	client := proto.NewUserClient(conn)

	resp, err := client.GetUserByMobile(ctx, &proto.MobileRequest{
		Mobile: passwordLoginForm.Mobile,
	})

	if err != nil {
		utils.HandleGrpcErrorToHttp(err, ctx)
	}

	// check pwd
	checkResp, err := client.CheckPassWord(ctx, &proto.PasswordCheckInfo{
		PassWord:          passwordLoginForm.PassWord,
		EncryptedPassword: resp.PassWord,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"password": "login failed",
		})
	}

	if checkResp.Success {
		// 生成token
		j := middleware.NewJWT()
		claims := jwtmodel.CustomClaims{
			ID:          uint(resp.Id),
			NickName:    resp.NickName,
			AuthorityId: uint(resp.Role),
			StandardClaims: jwt.StandardClaims{
				NotBefore: time.Now().Unix(),
				ExpiresAt: time.Now().Unix() + 60*60*24*30, // 30天
				Issuer:    "ws",
			},
		}
		token, err := j.CreateToken(claims)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "create token failed",
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"id":        resp.Id,
			"nick_name": resp.NickName,
			"token":     token,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"msg": "password error",
		})
	}
}
