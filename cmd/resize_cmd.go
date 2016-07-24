package cmd

import (
	"ImageScript/handler"
	"ImageScript/validity"

	"github.com/urfave/cli"
)

var resizeInfo = &validity.ImageResizeInfo{}

// Resize Icon 生成命令
var Resize = cli.Command{
	Name:      "resize",
	ShortName: "r",
	Usage:     "iOS Icon 生成命令",
	UsageText: "生成 iOS Icon 尺寸的图片文件",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "input, i",
			Usage:       "icon 文件路径",
			Destination: &resizeInfo.Input,
		},
	},
	Action: func(c *cli.Context) error {

		if err := resizeInfo.InitAndValidityOfImageResizeInfo(); err != nil {
			return err
		}

		if err := handler.IconResize(resizeInfo); err != nil {
			return err
		}
		return nil
	},
}
