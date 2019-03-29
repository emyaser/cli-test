package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	//实例化一个命令行程序
	app := cli.NewApp()
	//程序名称
	app.Name = "command test"
	//程序的用途描述
	app.Usage = "this is a command test for cli"
	//程序的版本号
	app.Version = "1.0.0"

	//设置多个命令处理函数
	app.Commands = []cli.Command{
		{
			//命令全称
			Name: "lang",
			//命令简写
			Aliases: []string{"l"},
			//命令详细描述
			Usage: "Setting language",
			//命令处理函数
			Action: func(c *cli.Context) {
				// 通过c.Args().First()获取命令行参数
				fmt.Printf("language=%v  \n", c.Args().First())
			},
		},
		{
			Name:    "encode",
			Aliases: []string{"e"},
			Usage:   "Setting encoding",
			Action: func(c *cli.Context) {
				fmt.Printf("encoding=%v \n", c.Args().First())
			},
		},
	}

	//启动
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
