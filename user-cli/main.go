package main

import (
	"context"
	"log"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"

	//pb "github.com/laizhenxing/laracom/user-service/proto/user"
)

func main() {

}

func callUserService() {
	// 初始化客户端服务,定义命令行参数标识
	service := micro.NewService(
		micro.Flags(
			&cli.StringFlag{Name: "name", Usage: "Your name"},
			&cli.StringFlag{Name: "email", Usage: "Your email"},
			&cli.StringFlag{Name: "password", Usage: "Your password"},
		),
	)

	// 远程服务客户端调用句柄
	client := pb.NewUserServiceClient("laracom.user.service", service.Client())

	// 运行客户端命令调用远程服务逻辑设置
	service.Init(
		micro.Action(func(ctx *cli.Context) error {
			name := ctx.String("name")
			email := ctx.String("email")
			password := ctx.String("password")

			log.Println("参数:", name, email, password)

			// 调用服务
			r, err := client.Create(context.TODO(), *pb.User{
				Name: name,
				Email: email,
				Password: password,
			})
			if err != nil {
				log.Printf("创建用户失败: %v\n", err)
			}
			log.Printf("创建用户成功: %s", r.User.Id)

			users, err := client.GetAll(context.Background(), &pb.UserRequest{})
			if err != nil {
				log.Printf("获取所有用户失败: %v\n", err)
			}
			for _, user := range users.Users {
				log.Println(user)
			}
			//os.Exit(0)
			return nil
		}),
	)

	if err := service.Run(); err != nil {
		log.Fatalf("用户客户端启动失败: %v\n", err)
	}
}
