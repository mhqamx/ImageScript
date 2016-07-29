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
	Usage:     "iOS icon 生成命令",
	UsageText: "生成 iOS icon 尺寸的图片文件",
	ArgsUsage: "input, output",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "input, i",
			Usage:       "icon 文件路径",
			Destination: &resizeInfo.Input,
		},
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "文件输出路径，若不指定或者错误，输出路径为脚本所在路径",
			Destination: &resizeInfo.Output,
		},
		cli.StringFlag{
			Name:        "perferred, p",
			Usage:       "输出文件类型 若不指定或者输入错误则输出三种设备的图片",
			Value:       "iPhone,iPad,iWatch",
			Destination: &resizeInfo.PerferredDevice,
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
