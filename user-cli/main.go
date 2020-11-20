package main

import (
	"context"
	"log"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"

	pb "github.com/laizhenxing/laracom/user-service/proto/user"
)

func main() {
	callUserService()
}

func callUserService() {
	// 初始化客户端服务,定义命令行参数标识
	service := micro.NewService(
		micro.Name("laracom.user.client"),
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

			// 调用用户注册服务服务
			r, err := client.Create(context.TODO(), &pb.User{
				Name: name,
				Email: email,
				Password: password,
			})
			if err != nil {
				log.Printf("创建用户失败: %v\n", err)
				return err
			}
			log.Printf("创建用户成功: %s", r.User.Id)

			// 调用用户认证服务
			var token *pb.Token
			token, err = client.Auth(context.TODO(), &pb.User{
				Email: email,
				Password: password,
			})
			if err != nil {
				log.Fatalf("用户登录失败:%v\n", err)
				return err
			}
			log.Printf("用户登录成功: %s\n", token.Token)

			// 调用用户验证服务
			token, err = client.ValidateToken(context.TODO(), token)
			if err != nil {
				log.Println("用户认证失败:", err)
				return err
			}
			log.Printf("用户认证成功：%t\n", token.Valid)

			// 调用获取所有用户的服务
			users, err := client.GetAll(context.Background(), &pb.UserRequest{})
			if err != nil {
				log.Printf("获取所有用户失败: %v\n", err)
				return err
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
