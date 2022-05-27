package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

// UploadController ...
type UploadController struct {
	beego.Controller
}

// Post ...
func (c *UploadController) Post() {
	f, h, err := c.GetFile("uploadname")
	if err != nil {
		fmt.Println("params err", err)
	}
	defer f.Close()
	c.SaveToFile("uploadname", "static/upload/"+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
}
