package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	//实例化一个命令程序
	app := cli.NewApp()
	//命令行名称
	app.Name = "cli test"
	//命令行功能描述
	app.Usage = "this is a base test for cli"
	//版本
	app.Version = "1.0.0"

	app.Action = func(c *cli.Context) error {
		fmt.Println("base test!")
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
