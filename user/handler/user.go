package handler

import (
	"context"

	"github.com/codeleongy/micro-market/user/domain/model"
	"github.com/codeleongy/micro-market/user/domain/service"
	pb "github.com/codeleongy/micro-market/user/proto/user"
)

type User struct {
	UserDataService service.IUserDataService
}

// 注册
func (u *User) Register(ctx context.Context, req *pb.UserRegisterRequest, res *pb.UserRegisterResponse) error {
	userRegister := &model.User{
		UserName:     req.UserName,
		FirstName:    req.FirstName,
		HashPassword: req.Pwd,
	}

	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}

	res.Message = "添加成功"
	return nil
}

// 登录
func (u *User) Login(ctx context.Context, req *pb.UserLoginRequest, res *pb.UserLoginResponse) error {
	isOK, err := u.UserDataService.CheckPwd(req.UserName, req.Pwd)
	if err != nil {
		return err
	}

	res.IsSuccess = isOK
	return nil
}

// 查询用户信息
func (u *User) GetUserInfo(ctx context.Context, req *pb.UserInfoRequest, res *pb.UserInfoResponse) error {
	userInfor, err := u.UserDataService.FindUserByName(req.UserName)
	if err != nil {
		return err
	}
	UserForRes(userInfor, res)

	return nil
}

// 类型转换
func UserForRes(userModel *model.User, res *pb.UserInfoResponse) {
	res.UserId = userModel.ID
	res.UserName = userModel.UserName
	res.FirstName = userModel.FirstName
}
