package controllers

import (
	"backend/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

// InfoController ...
type InfoController struct {
	beego.Controller
}

// Get ...
// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (c *InfoController) Get() {
	infos := models.GetAllList()
	c.Data["json"] = infos
	c.ServeJSON()
}

// Post ...
// @Title create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (c *InfoController) Post() {
	var info models.Info
	fmt.Println("begin----------------")
	fmt.Println(string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &info)
	err := models.AddInfo(info)
	if err != nil {
		c.Data["json"] = map[string]string{"msg": "failed"}
		fmt.Println("failed----------------", fmt.Sprintf("%s", err))
	} else {
		c.Data["json"] = map[string]string{"msg": "success"}
		fmt.Println("success----------------")
	}
	c.ServeJSON()
}

// GetListByParams ...
// @Title Get
// @Description find list by pageIndex, pageSize
// @Param	useragent		path 	string	true		"the pageIndex, pageSize you want to get"
// @router /:page [get]
func (c *InfoController) GetListByParams() {
	pageIndex := c.Ctx.Input.Query("pageIndex")
	pageSize := c.Ctx.Input.Query("pageSize")
	infos := models.GetListByParams(pageIndex, pageSize)
	c.Data["json"] = infos
	c.ServeJSON()
}
