package conf

import (
	"fmt"

	"github.com/joho/godotenv"
)

// 初始化配置项
func Init()  {
	// 从本地读取环境变量
	err := godotenv.Load()
	if err != nil {
		fmt.Println("配置加载失败:", err)
	}
}
