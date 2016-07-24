package handler

import (
	"fmt"
	"image/png"
	"os"

	"ImageScript/validity"

	"github.com/nfnt/resize"
)

// IconInfo iOS 设备所需Icon信息
type IconInfo struct {
	DeviceName  string
	BaseWidth   float32
	BasegHeight float32
	Multiply    uint
}

var iconInfos = []*IconInfo{
	&IconInfo{
		DeviceName:  "iPhone Spotlight(iOS 5-6) Settings(iOS 5-9)",
		BaseWidth:   29.0,
		BasegHeight: 29.0,
		Multiply:    2,
	},
	&IconInfo{
		DeviceName:  "iPhone Spotlight(iOS 5-6) Settings(iOS 5-9)",
		BaseWidth:   29.0,
		BasegHeight: 29.0,
		Multiply:    3,
	},
	&IconInfo{
		DeviceName:  "iPhone Spotlight(iOS 7-9)",
		BaseWidth:   40.0,
		BasegHeight: 40.0,
		Multiply:    2,
	},
	&IconInfo{
		DeviceName:  "iPhone Spotlight(iOS 7-9)",
		BaseWidth:   40.0,
		BasegHeight: 40.0,
		Multiply:    3,
	},
	&IconInfo{
		DeviceName:  "iPhone App(iOS 7-9)",
		BaseWidth:   60.0,
		BasegHeight: 60.0,
		Multiply:    2,
	},
	&IconInfo{
		DeviceName:  "iPhone App(iOS 7-9)",
		BaseWidth:   60.0,
		BasegHeight: 60.0,
		Multiply:    3,
	},
	&IconInfo{
		DeviceName:  "iPad Setting(iOS 5-9)",
		BaseWidth:   29.0,
		BasegHeight: 29.0,
		Multiply:    1,
	},
	&IconInfo{
		DeviceName:  "iPad Setting(iOS 5-9)",
		BaseWidth:   29.0,
		BasegHeight: 29.0,
		Multiply:    2,
	},
	&IconInfo{
		DeviceName:  "iPad Spotlight(iOS 7-9)",
		BaseWidth:   40.0,
		BasegHeight: 40.0,
		Multiply:    1,
	},
	&IconInfo{
		DeviceName:  "iPad Spotlight(iOS 7-9)",
		BaseWidth:   40.0,
		BasegHeight: 40.0,
		Multiply:    2,
	},
	&IconInfo{
		DeviceName:  "iPad App(iOS 7-9)",
		BaseWidth:   76.0,
		BasegHeight: 76.0,
		Multiply:    1,
	},
	&IconInfo{
		DeviceName:  "iPad App(iOS 7-9)",
		BaseWidth:   76.0,
		BasegHeight: 76.0,
		Multiply:    2,
	},
	&IconInfo{
		DeviceName:  "iPad Pro App(iOS 9)",
		BaseWidth:   83.5,
		BasegHeight: 83.5,
		Multiply:    2,
	},
	&IconInfo{
		DeviceName:  "Apple Watch Notification Center 38mm",
		BaseWidth:   24.0,
		BasegHeight: 24.0,
		Multiply:    2,
	},
	&IconInfo{
		DeviceName:  "Apple Watch Notification Center 42mm",
		BaseWidth:   27.5,
		BasegHeight: 27.5,
		Multiply:    2,
	},
	&IconInfo{
		DeviceName:  "Apple Watch Companion Settings",
		BaseWidth:   29.0,
		BasegHeight: 29.0,
		Multiply:    2,
	},
	&IconInfo{
		DeviceName:  "Apple Watch Companion Settings",
		BaseWidth:   29.0,
		BasegHeight: 29.0,
		Multiply:    3,
	},
	&IconInfo{
		DeviceName:  "Apple Watch Home Screen(All)",
		BaseWidth:   40.0,
		BasegHeight: 40.0,
		Multiply:    2,
	},
	&IconInfo{
		DeviceName:  "Apple Watch Short Look 38mm",
		BaseWidth:   86.0,
		BasegHeight: 86.0,
		Multiply:    2,
	},
	&IconInfo{
		DeviceName:  "Apple Watch Short Look 42mm",
		BaseWidth:   98.0,
		BasegHeight: 98.0,
		Multiply:    2,
	},
}

// UltimateImageSize 计算resize之后图片实际大小
func (info *IconInfo) UltimateImageSize() (width, height uint) {
	return uint(info.BaseWidth * float32(info.Multiply)), uint(info.BasegHeight * float32(info.Multiply))
}

// AssembleImageName 拼接Image的名称
func (info *IconInfo) AssembleImageName() string {

	width, height := info.UltimateImageSize()
	imageName := info.DeviceName
	return fmt.Sprintf("%s_%dx%d_@%dx.png", imageName, width, height, info.Multiply)
}

// IconResize 转换输入的Icon
func IconResize(imageResizeInfo *validity.ImageResizeInfo) error {

	file, err := os.Open(imageResizeInfo.Input)

	if err != nil {
		return fmt.Errorf("图片文件打开失败\n错误信息：%s", err.Error())
	}

	img, err := png.Decode(file)

	for _, info := range iconInfos {

		w, h := info.UltimateImageSize()

		fn := info.AssembleImageName()

		m := resize.Resize(w, h, img, resize.Lanczos3)

		out, err := os.Create(fn)

		if err != nil {
			return fmt.Errorf("创建图片出错\n错误信息：%s", err.Error())
		}

		defer out.Close()

		png.Encode(out, m)
	}

	return nil
}
