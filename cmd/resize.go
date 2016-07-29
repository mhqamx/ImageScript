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

		//		cli.StringFlag{
		//			Name:  "iph",
		//			Value: "true",
		//			Usage: "是否输出iPhone图片",
		//		},
		//		cli.StringFlag{
		//			Name:  "ipa",
		//			Value: "true",
		//			Usage: "是否输出iPad图片",
		//		},
		//		cli.StringFlag{
		//			Name:  "iwa",
		//			Value: "true",
		//			Usage: "是否输出iWatch图片",
		//		},

		//你说用上边的写法还是下边的写法  上面的不用写等号 下边的还要输等号...
		// 建议添加一个 PerferredDevice 字段到Info中, 还有最好是发Pull request或者Merge request吗? 这个我根本看不到啊
		// GO中处理字符串还是挺方便的
		cli.BoolTFlag{
			Name:        "iph",
			Usage:       "是否输出iPhone图片",
			Destination: &resizeInfo.IsHaveiPhone,
		},
		cli.BoolTFlag{
			Name:        "ipa",
			Usage:       "是否输出iPad图片",
			Destination: &resizeInfo.IsHaveiPad,
		},
		cli.BoolTFlag{
			Name:        "iwa",
			Usage:       "是否输出iWatch图片",
			Destination: &resizeInfo.IsHaveiWatch,
		},
	},
	Action: func(c *cli.Context) error {

		//		resizeInfo.IsHaveiPhone = c.BoolT("iph")
		//		resizeInfo.IsHaveiPad = c.BoolT("ipa")
		//		resizeInfo.IsHaveiWatch = c.BoolT("iwa")

		if err := resizeInfo.InitAndValidityOfImageResizeInfo(); err != nil {
			return err
		}

		if err := handler.IconResize(resizeInfo); err != nil {
			return err
		}
		return nil
	},
}
