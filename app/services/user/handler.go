package main

import (
	"context"
	"fmt"

	rpc_user "github.com/bitdance-panic/gobuy/app/rpc/kitex_gen/user"
	"github.com/bitdance-panic/gobuy/app/services/user/biz/bll"
)

var ub *bll.UserBLL

func init() {
	ub = bll.NewUserBLL()
}

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *rpc_user.RegisterReq) (resp *rpc_user.RegisterResp, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *rpc_user.LoginReq) (resp *rpc_user.LoginResp, err error) {
	fmt.Printf("Login called, req: %+v\n", *req)
	// 做validate...
	resp, err = ub.Login(ctx, req)
	return
}
