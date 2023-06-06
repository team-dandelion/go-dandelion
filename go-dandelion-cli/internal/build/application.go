package build

import (
	"github.com/gly-hub/toolbox/file"
	"path"
)

// 构建应用基础架构

// BuildApplication 创建应用
func BuildApplication(app string) {
	pwd, err := file.GetPwd()
	if err != nil {
		return
	}
	appDir := path.Join(pwd, app)
	// 创建应用文件夹
	if err = file.CreateDir(appDir); err != nil {
		return
	}
}
