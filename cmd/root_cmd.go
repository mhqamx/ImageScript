package cmd

import (
	"os"

	"ImageScript/info"

	"github.com/urfave/cli"
)

// Root Root命令
var Root = cli.NewApp()

// Execute 执行Root命令
func Execute() {
	if err := Root.Run(os.Args); err != nil {
		panic(err.Error())
	}
}

func init() {
	cli.AppHelpTemplate = info.AppHelpTemplate
	Root.Name = "ImageSpitit"
	Root.Usage = "iOS开发图片处理工具"
	Root.Version = info.Version
	Root.Author = info.Author
	Root.Email = info.Email
	Root.Commands = []cli.Command{
		Resize,
	}
}
