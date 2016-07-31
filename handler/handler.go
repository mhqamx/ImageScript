package handler

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"sync"

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

var iPhoneInfos = []*IconInfo{
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
}

var iPadInfos = []*IconInfo{
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
}

var appleWatchInfos = []*IconInfo{
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

// IconResize 转换输入的Icon
func IconResize(imageResizeInfo *validity.ImageResizeInfo) error {

	file, err := os.Open(imageResizeInfo.Input)

	if err != nil {
		return fmt.Errorf("图片文件打开失败\n错误信息：%s", err.Error())
	}

	img, err := png.Decode(file)

	if err != nil {
		return fmt.Errorf("图片文件解析失败\n错误信息：%s", err.Error())
	}

	wg := &sync.WaitGroup{}

	switch imageResizeInfo.PreferenceDevice {
	case "phone":
		render(img, iPhoneInfos, imageResizeInfo, "iPhone icons", wg)
	case "pad":
		render(img, iPadInfos, imageResizeInfo, "iPad icons", wg)
	case "watch":
		render(img, appleWatchInfos, imageResizeInfo, "Apple Watch icons", wg)
	case "all":
		render(img, iPhoneInfos, imageResizeInfo, "iPhone icons", wg)
		render(img, iPadInfos, imageResizeInfo, "iPad icons", wg)
		render(img, appleWatchInfos, imageResizeInfo, "Apple Watch icons", wg)
	default:
		return fmt.Errorf("发现未知设备信息")
	}

	wg.Wait()

	return nil
}

func render(img image.Image, info []*IconInfo, imageResizeInfo *validity.ImageResizeInfo, folder string, wg *sync.WaitGroup) {
	wg.Add(1)
	go func(img image.Image, infos []*IconInfo) {
		defer wg.Done()
		assembleFilePathAndDistribute(img, infos, imageResizeInfo.Output, fmt.Sprintf("/%s", folder))
	}(img, info)
}

func (info *IconInfo) ultimateImageSize() (width, height uint) {
	return uint(info.BaseWidth * float32(info.Multiply)), uint(info.BasegHeight * float32(info.Multiply))
}

func (info *IconInfo) assembleImageName() string {

	width, height := info.ultimateImageSize()
	imageName := info.DeviceName
	return fmt.Sprintf("%s_%dx%d_@%dx.png", imageName, width, height, info.Multiply)
}

func assembleFilePathAndDistribute(img image.Image, infos []*IconInfo, outputPath string, deviceFolder string) {

	fp := fmt.Sprintf("%s%s", outputPath, deviceFolder)
	if err := os.MkdirAll(fp, 0755); err != nil {
		panic(fmt.Sprintf("创建文件夹失败: %s", err.Error()))
	}
	for _, info := range infos {
		outputImage(img, fp, info)
	}
}

func outputImage(img image.Image, fp string, info *IconInfo) {

	w, h := info.ultimateImageSize()
	fn := info.assembleImageName()
	m := resize.Resize(w, h, img, resize.Lanczos3)
	var path string
	if fp != "" {
		path = fmt.Sprintf("%s/%s", fp, fn)
	}
	out, err := os.Create(path)
	if err != nil {
		fmt.Println("构建失败：", path, err.Error())
	}
	defer out.Close()
	png.Encode(out, m)
}
