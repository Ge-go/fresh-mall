package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"go.uber.org/zap"
	"strings"
	"time"

	"github.com/anaskhan96/go-password-encoder"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"

	"mall_srvs/goods_srv/global"
	"mall_srvs/goods_srv/model"
	"mall_srvs/goods_srv/proto"
)

type UserServer struct {
	proto.UnsafeUserServer
}

func ModelToUserInfoResponse(user model.User) *proto.UserInfoResponse {
	userInfoRsp := &proto.UserInfoResponse{
		Id:       user.ID,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
		Mobile:   user.Mobile,
		PassWord: user.Password,
	}
	if user.Birthday != nil {
		userInfoRsp.BirthDay = uint64(user.Birthday.Unix())
	}

	return userInfoRsp
}

// Paginate gorm内置分页
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// GetUserList user manager获取用户列
func (u *UserServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	// todo 实际企业级开发到底如何优化这里
	var users []model.User
	res := global.DB.WithContext(ctx).Find(&users)
	if res.Error != nil {
		zap.S().Errorw("get user total err", "msg", res.Error.Error())
		return nil, res.Error
	}

	rsp := &proto.UserListResponse{}
	rsp.Total = int32(res.RowsAffected)

	res = global.DB.WithContext(ctx).Scopes(Paginate(int(req.Pn), int(req.PSize))).Order("id").Find(&users)
	if res.Error != nil {
		zap.S().Errorw("get user list err", "msg", res.Error.Error())
		return nil, res.Error
	}

	for _, user := range users {
		userInfoRsp := ModelToUserInfoResponse(user)
		rsp.Data = append(rsp.Data, userInfoRsp)
	}

	return rsp, nil
}

// GetUserByMobile 通过Mobile 查询用户
func (u *UserServer) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	res := global.DB.WithContext(ctx).Where(&model.User{Mobile: req.Mobile}).First(&user)
	if res.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "not found user by mobile")
	}
	if res.Error != nil {
		zap.S().Errorw("get user by mobile err", "msg", res.Error.Error())
		return nil, res.Error
	}

	userInfoRsp := ModelToUserInfoResponse(user)
	return userInfoRsp, nil
}

// GetUserById 通过id查找用户
func (u *UserServer) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	res := global.DB.WithContext(ctx).First(&user, req.Id)
	if res.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "not found user by mobile")
	}
	if res.Error != nil {
		zap.S().Errorw("get user by id err", "msg", res.Error.Error())
		return nil, res.Error
	}

	userInfoRsp := ModelToUserInfoResponse(user)
	return userInfoRsp, nil
}

// CreateUser 新建用户
func (u *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	var user model.User
	res := global.DB.WithContext(ctx).Where(&model.User{Mobile: req.Mobile}).First(&user)
	if res.RowsAffected == 1 {
		return nil, status.Error(codes.AlreadyExists, "mobile is already exists")
	}
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		zap.S().Errorw("find user by create user err", "msg", res.Error.Error())
		return nil, status.Error(codes.Internal, res.Error.Error())
	}

	user.Mobile = req.Mobile
	user.NickName = req.NickName

	//密码加密
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodePwd := password.Encode(req.PassWord, options)
	user.Password = fmt.Sprintf("$pdkdf2-sha512$%s$%s", salt, encodePwd)

	res = global.DB.WithContext(ctx).Create(&user)
	if res.Error != nil {
		zap.S().Errorw("create user err", "msg", res.Error.Error())
		return nil, status.Error(codes.Internal, res.Error.Error())
	}

	userInfoRsp := ModelToUserInfoResponse(user)
	return userInfoRsp, nil
}

func (u *UserServer) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (*emptypb.Empty, error) {
	var user model.User
	res := global.DB.WithContext(ctx).First(&user, req.Id)
	if res.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "not found user by id")
	}
	if res.Error != nil {
		zap.S().Errorw("get user by id err", "msg", res.Error.Error())
		return nil, res.Error
	}

	birthday := time.Unix(int64(req.BirthDay), 0)
	user.Birthday = &birthday
	user.Gender = req.Gender
	user.NickName = req.NickName

	res = global.DB.WithContext(ctx).Save(&user)
	if res.Error != nil {
		zap.S().Errorw("update user err", "msg", res.Error.Error())
		return nil, status.Error(codes.Internal, res.Error.Error())
	}

	return &emptypb.Empty{}, nil
}

func (u *UserServer) CheckPassWord(ctx context.Context, req *proto.PasswordCheckInfo) (*proto.CheckResponse, error) {
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	passwordInfo := strings.Split(req.EncryptedPassword, "$")

	if len(passwordInfo) < 4 { //做一个保护,不然服务器会发生崩溃
		return nil, status.Error(codes.InvalidArgument, "invalid encryptedPassword")
	}
	check := password.Verify(req.PassWord, passwordInfo[2], passwordInfo[3], options)
	return &proto.CheckResponse{
		Success: check,
	}, nil
}
