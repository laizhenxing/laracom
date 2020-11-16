// user service entry
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/micro/go-micro/v2"

	"github.com/laizhenxing/laracom/user-service/conf"
	"github.com/laizhenxing/laracom/user-service/db"
	"github.com/laizhenxing/laracom/user-service/handler"
	pb "github.com/laizhenxing/laracom/user-service/proto/user"
	"github.com/laizhenxing/laracom/user-service/repo"
)

func main() {
	// 初始化配置
	conf.Init()

	// 创建数据库连接
	dbConn, err := db.CreateDBConnection()
	if err != nil {
		log.Fatalf("连接数据库失败: %v\n", err)
		os.Exit(1)
	}
	defer dbConn.Close()

	// 数据库迁移
	dbConn.AutoMigrate(&pb.User{})

	// 初始化 Repo
	rep := &repo.UserRepository{dbConn}

	// 创建用户服务
	srv := micro.NewService(
		micro.Name("laracom.user.service"),
		micro.Version("latest"), // 新增加接口版本参数
	)
	// 解析命令行参数
	srv.Init()
	// 注册服务处理器
	pb.RegisterUserServiceHandler(srv.Server(), &handler.UserService{rep})
	// 运行服务
	if err := srv.Run(); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
