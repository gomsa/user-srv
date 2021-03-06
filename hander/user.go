package hander

import (
	"context"
	"fmt"

	pb "github.com/gomsa/user/proto/user"
	"github.com/gomsa/user/service"

	"golang.org/x/crypto/bcrypt"
)

// User 用户结构
type User struct {
	Repo service.URepository
}

// Exist 用户是否存在
func (srv *User) Exist(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	res.Valid = srv.Repo.Exist(req.User)
	return err
}

// List 获取所有用户
func (srv *User) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	users, err := srv.Repo.List(req.ListQuery)
	total, err := srv.Repo.Total(req.ListQuery)
	if err != nil {
		return err
	}
	res.Total = total
	res.Users = users
	return err
}

// Get 获取用户
func (srv *User) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	user, err := srv.Repo.Get(req.User)
	if err != nil {
		return err
	}
	res.User = user
	return err
}

// Create 创建用户
func (srv *User) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.User.Password = string(hashedPass)
	user, err := srv.Repo.Create(req.User)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("创建用户失败")
	}
	res.User = user
	res.Valid = true
	return err
}

// Update 更新用户
func (srv *User) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 密码不为空时才可以修改密码
	if req.User.Password != "" {
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.User.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		req.User.Password = string(hashedPass)
	}
	valid, err := srv.Repo.Update(req.User)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新用户失败")
	}
	res.Valid = valid
	return err
}

// Delete 删除用户用户
func (srv *User) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req.User)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("删除用户失败")
	}
	res.Valid = valid
	return err
}
