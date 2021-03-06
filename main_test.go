package main

import (
	"context"
	"testing"

	_ "github.com/gomsa/user/database/migrations"
	db "github.com/gomsa/user/providers/database"

	"github.com/gomsa/user/hander"
	authPB "github.com/gomsa/user/proto/auth"
	frontPermitPB "github.com/gomsa/user/proto/frontPermit"
	permissionPB "github.com/gomsa/user/proto/permission"
	userPB "github.com/gomsa/user/proto/user"
	"github.com/gomsa/user/service"
)

func TestFrontPermitUpdateOrCreate(t *testing.T) {
	req := &frontPermitPB.Request{
		FrontPermit: &frontPermitPB.FrontPermit{
			App: "ui", Service: "role", Method: "permission", Name: "角色权限", Description: "管理角色权限。",
		},
	}
	repo := &service.FrontPermitRepository{db.DB}
	h := hander.FrontPermit{repo}
	res := &frontPermitPB.Response{}
	err := h.UpdateOrCreate(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}
func TestPermissionsUpdateOrCreate(t *testing.T) {
	req := &permissionPB.Request{
		Permission: &permissionPB.Permission{
			Service: "user-api", Method: "Auth.Auth1", Name: "用户授权3", Description: "用户登录授权返回 token 权限。",
		},
	}
	repo := &service.PermissionRepository{db.DB}
	h := hander.Permission{repo}
	res := &permissionPB.Response{}
	err := h.UpdateOrCreate(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}
func TestUserCreate(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	h := hander.User{repo}
	req := &userPB.Request{
		User: &userPB.User{
			Username: `bvbv0115`,
			Password: `123456`,
			Mobile:   `13953186115`,
			Email:    `bvbv0a115@qq.com`,
			Name:     `bvbv0111`,
			Avatar:   `https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif`,
			Origin:   `user`,
		},
	}
	res := &userPB.Response{}
	err := h.Create(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}

func TestUserIsExist(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	h := hander.User{repo}
	req := &userPB.Request{
		User: &userPB.User{
			Username: `bvbv0115`,
			Mobile:   `13953186115`,
			Email:    `bvbv0a1@qq.com`,
		},
	}
	res := &userPB.Response{}
	err := h.Exist(context.TODO(), req, res)
	// fmt.Println(req, res.Valid, err)
	t.Log(req, res, err)
}
func TestUserGet(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	h := hander.User{repo}
	req := &userPB.Request{
		User: &userPB.User{
			Username: `bvbv0115`,
		},
	}
	res := &userPB.Response{}
	err := h.Get(context.TODO(), req, res)
	// fmt.Println("UserGet", res, err)
	t.Log(req, res, err)
}

func TestUserList(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	h := hander.User{repo}
	req := &userPB.Request{
		ListQuery: &userPB.ListQuery{
			Limit: 20,
			Page:  1,
			Sort:  "created_at desc",
		},
	}
	res := &userPB.Response{}
	err := h.List(context.TODO(), req, res)
	// fmt.Println("UserList", res, err)
	t.Log(req, res, err)
}
func TestUserUpdate(t *testing.T) {
	// repo := &service.UserRepository{db.DB}
	// h := hander.User{repo}
	// req := &userPB.Request{
	// 	User: &userPB.User{
	// 		Id:       `8cd1d57f-6f53-49e4-b751-96eefc4f4b20`,
	// 		Username: `bvbv0111`,
	// 		Name:     `newbvbv`,
	// 	},
	// }
	// res := &userPB.Response{}
	// err := h.Update(context.TODO(), req, res)
	// fmt.Println("UserUpdate", req, res, err)
	// t.Log(req, res, err)
}
func TestUserDelete(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	h := hander.User{repo}
	req := &userPB.Request{
		User: &userPB.User{
			Id: `8cd1d57f-6f53-49e4-b751-96eefc4f4b20`,
		},
	}
	res := &userPB.Response{}
	err := h.Delete(context.TODO(), req, res)
	// fmt.Println("UserDelete", req, res, err)
	t.Log(req, res, err)
}

// Auth
func TestAuth(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	token := &service.TokenService{}
	h := hander.Auth{token, repo}
	req := &authPB.Request{
		User: &authPB.User{
			Username: `bvbv011`,
			Password: `123456`,
		},
	}
	res := &authPB.Response{}
	err := h.Auth(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}

func TestAuthById(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	token := &service.TokenService{}
	h := hander.Auth{token, repo}
	req := &authPB.Request{
		User: &authPB.User{
			Id: `c0a83b24-c01c-4601-a1c6-17e3c1864c5a`,
		},
	}
	res := &authPB.Response{}
	err := h.AuthById(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}

func TestValidateToken(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	token := &service.TokenService{}
	h := hander.Auth{token, repo}
	req := &authPB.Request{
		Token: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiYzBhODNiMjQtYzAxYy00NjAxLWExYzYtMTdlM2MxODY0YzVhIn0sImV4cCI6MTU3MDM0OTc2MCwiaXNzIjoidXNlciJ9.Y3l55bE3StZL66RPbrTk8zVgUZBll0Pc6yV6ljb22k4`,
	}
	res := &authPB.Response{}
	err := h.ValidateToken(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}
