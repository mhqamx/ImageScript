package validity

import (
	"fmt"
	"path/filepath"
	"strings"
)

// ImageResizeInfo 命令信息结构
type ImageResizeInfo struct {
	Input         string
	Dir           string
	File          string
	FileName      string
	FileExtension string
}

// InitAndValidityOfImageResizeInfo 检查输入命令合法性
func (info *ImageResizeInfo) InitAndValidityOfImageResizeInfo() error {

	if info.Input == "" {
		return fmt.Errorf("Icon文件路径为空")
	}

	p, err := filepath.Abs(info.Input)

	if err != nil {
		return fmt.Errorf("Icon文件路径错误\n错误信息：%s", info.Input)
	}

	_, file := filepath.Split(p)

	if file == "" {
		return fmt.Errorf("Icon文件路径错误\n错误信息：%s", info.Input)
	}

	components := strings.Split(file, ".")

	if len(components) < 2 {
		return fmt.Errorf("路径错误\n错误信息：%s不是一个*.png图片的路径", info.Input)
	}

	if components[1] != "png" {
		return fmt.Errorf("图片必须为*.png格式")
	}

	info.Input = p
	info.File = file
	info.FileName = components[0]
	info.FileExtension = components[1]

	return nil
}
