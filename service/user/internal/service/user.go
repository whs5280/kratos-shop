package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "user/api/user/v1"
	"user/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc  *biz.UserUsecase
	log *log.Helper
}

// NewUserService new a greeter service.
func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (u *UserService) CreateUser(ctx context.Context, req *pb.CreateUserInfo) (*pb.UserInfoResponse, error) {
	user, err := u.uc.Create(ctx, &biz.User{
		Mobile:   req.Mobile,
		Password: req.Password,
		Nickname: req.NickName,
	})
	if err != nil {
		return nil, err
	}

	userInfoRsp := pb.UserInfoResponse{
		Id:       user.Id,
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.Nickname,
		Gender:   user.Gender,
		Role:     int32(user.Role),
		Birthday: user.Birthday,
	}
	return &userInfoRsp, nil
}
