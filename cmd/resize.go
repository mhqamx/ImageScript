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
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "device, d",
			Usage:       "期望生成iocn的设备选项",
			Value:       "phone, pad, watch, all",
			Destination: &resizeInfo.PreferenceDevice,
		},
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
