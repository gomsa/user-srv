package main

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/gomsa/user-srv/database/migrations"
	db "github.com/gomsa/user-srv/providers/database"

	"github.com/gomsa/user-srv/hander"
	authPB "github.com/gomsa/user-srv/proto/auth"
	userPB "github.com/gomsa/user-srv/proto/user"
	"github.com/gomsa/user-srv/service"
)

func TestUserCreate(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	h := hander.User{repo}
	req := &userPB.User{
		Username: `bvbv0111`,
		Password: `123456`,
		Mobile:   `13953186114`,
		Email:    `bvbv0a1@qq.com`,
	}
	res := &userPB.Response{}
	err := h.Create(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}

func TestUserIsExist(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	h := hander.User{repo}
	req := &userPB.User{
		Username: `bvbv0111`,
		Mobile:   `13953186114`,
		Email:    `bvbv0a1@qq.com`,
		Origin:   `user-srv`,
	}
	res := &userPB.Response{}
	err := h.Create(context.TODO(), req, res)
	// fmt.Println(req, res.Valid, err)
	t.Log(req, res, err)
}
func TestUserGet(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	h := hander.User{repo}
	req := &userPB.User{
		Username: `bvbv011`,
	}
	res := &userPB.Response{}
	err := h.Get(context.TODO(), req, res)
	// fmt.Println(res, err)
	t.Log(req, res, err)
}
func TestUserGetAll(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	h := hander.User{repo}
	req := &userPB.Request{}
	res := &userPB.Response{}
	err := h.GetAll(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}

// Auth
func TestAuth(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	token := &service.TokenService{}
	h := hander.Auth{token, repo}
	req := &authPB.User{
		Username: `bvbv0111`,
		Password: `123456`,
	}
	res := &authPB.Token{}
	err := h.Auth(context.TODO(), req, res)
	fmt.Println(req, res, err)
	t.Log(req, res, err)
}

func TestAuthById(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	token := &service.TokenService{}
	h := hander.Auth{token, repo}
	req := &authPB.User{
		Id: `93267c95-ef44-4e49-b865-0ad6c84c646c`,
	}
	res := &authPB.Token{}
	err := h.AuthById(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}

func TestValidateToken(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	token := &service.TokenService{}
	h := hander.Auth{token, repo}
	req := &authPB.Request{
		Token: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiOTMyNjdjOTUtZWY0NC00ZTQ5LWI4NjUtMGFkNmM4NGM2NDZjIiwibmFtZSI6ImJ2YnYwMTEiLCJtb2JpbGUiOiIxMzk1NDM4NjExNCIsImVtYWlsIjoiYnZidjBhMUBxcS5jb20iLCJwYXNzd29yZCI6IiQyYSQxMCRFM3VDdzRpa3JNaXhaSjhWd21tYmRlMERyUlA3WDNaM0xOMEdMZ2Z5SWxNTWx0MGlZazZDRyIsIm9yaWdpbiI6InVzZXItc2VydmljZSIsImNyZWF0ZWRfYXQiOiIyMDE5LTA1LTA2VDE2OjI1OjMyKzA4OjAwIiwidXBkYXRlZF9hdCI6IjIwMTktMDUtMDZUMTY6MjU6MzIrMDg6MDAifSwiZXhwIjoxNTU3NDUyMDEzLCJpc3MiOiJ1c2VyLXNlcnZpY2UifQ.kQKnSW-qDGyVf3XahRFkkTx5wP_PyIyM_RGqiBK4pfI`,
	}
	res := &authPB.Token{}
	err := h.ValidateToken(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}
