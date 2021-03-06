package main

import (
	// 公共引入
	k8s "github.com/micro/examples/kubernetes/go/micro"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	// 执行数据迁移
	_ "github.com/gomsa/user/database/migrations"

	"github.com/gomsa/user/hander"
	authPB "github.com/gomsa/user/proto/auth"
	casbinPB "github.com/gomsa/user/proto/casbin"
	frontPermitPB "github.com/gomsa/user/proto/frontPermit"
	permissionPB "github.com/gomsa/user/proto/permission"
	rolePB "github.com/gomsa/user/proto/role"
	userPB "github.com/gomsa/user/proto/user"
	"github.com/gomsa/user/providers/casbin"
	db "github.com/gomsa/user/providers/database"
	"github.com/gomsa/user/service"
)

func main() {
	srv := k8s.NewService(
		micro.Name(Conf.Service),
		micro.Version(Conf.Version),
	)
	srv.Init()

	// 用户服务实现
	repo := &service.UserRepository{db.DB}
	userPB.RegisterUsersHandler(srv.Server(), &hander.User{repo})

	// token 服务实现
	token := &service.TokenService{}
	authPB.RegisterAuthHandler(srv.Server(), &hander.Auth{token, repo})

	// 前端权限服务实现
	fprepo := &service.FrontPermitRepository{db.DB}
	frontPermitPB.RegisterFrontPermitsHandler(srv.Server(), &hander.FrontPermit{fprepo})

	// 权限服务实现
	prepo := &service.PermissionRepository{db.DB}
	permissionPB.RegisterPermissionsHandler(srv.Server(), &hander.Permission{prepo})

	// 角色服务实现
	rrepo := &service.RoleRepository{db.DB}
	rolePB.RegisterRolesHandler(srv.Server(), &hander.Role{rrepo})

	// 权限管理服务实现
	casbinPB.RegisterCasbinHandler(srv.Server(), &hander.Casbin{casbin.Enforcer})
	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}
