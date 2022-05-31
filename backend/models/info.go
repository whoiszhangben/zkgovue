package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"

	// mysql ...
	_ "github.com/go-sql-driver/mysql"
)

var (
	// AgentList ...
	AgentList map[string]*Info
)

// Info ...
type Info struct {
	Id         int    `json:"id"`
	Appversion string `json:"appversion"`
	Appname    string `json:"appname"`
	Language   string `json:"language"`
	Platform   string `json:"platform"`
	Useragent  string `json:"useragent"`
}

// PackInfo ...
type PackInfo struct {
	Code   int    `json:"id"`
	Msg    string `json:"msg"`
	Data   []Info `json:"data"`
	Length int64  `json:"total"`
}

// TableName ...
func (u *Info) TableName() string {
	return "info"
}

func init() {
	//注册模型（库表）
	orm.RegisterModel(new(Info))
}

// AddInfo ...
func AddInfo(info Info) error {
	var o orm.Ormer
	o = orm.NewOrm()
	o.Using("ben")
	_, err := o.Insert(&info)
	return err
}

// GetAllUsers ...
func GetAllUsers() PackInfo {
	var o orm.Ormer
	o = orm.NewOrm()
	o.Using("ben")
	var rtos []Info
	var sql string = "select id, appname, language, platform from info"
	_, err := o.Raw(sql).QueryRows(&rtos)
	if err != nil {
		fmt.Printf("sql find info failed! %s\n", err)
	}
	p := PackInfo{}
	p.Code = 200
	p.Msg = "查询成功"
	p.Length = int64(len(rtos))
	p.Data = rtos
	return p
}

// GetAllList ...
func GetAllList() []Info {
	var o orm.Ormer
	o = orm.NewOrm()
	o.Using("ben")
	var rtos []Info
	var sql string = "select id, appname, language, platform from info"
	_, err := o.Raw(sql).QueryRows(&rtos)
	if err != nil {
		fmt.Printf("sql find info failed! %s\n", err)
	}
	return rtos
}

// GetListByParams ...
func GetListByParams(pageIndex string, pageSize string) PackInfo {
	var o orm.Ormer
	o = orm.NewOrm()
	o.Using("ben")
	intIndex, err := strconv.Atoi(pageIndex)
	if err != nil {
		p := PackInfo{}
		p.Code = 500
		p.Msg = "分页参数不正确"
		p.Length = 0
		p.Data = nil
		return p
	}
	intSize, err := strconv.Atoi(pageSize)
	if err != nil {
		p := PackInfo{}
		p.Code = 500
		p.Msg = "分页参数不正确"
		p.Length = 0
		p.Data = nil
		return p
	}
	var rtos []Info
	tableName := new(Info)
	qs := o.QueryTable(tableName)
	cnt, err := qs.Count()
	if err != nil {
		fmt.Printf("sql get count %s\n", err)
	}
	_, err = qs.Offset(intIndex * intSize).Limit(intSize).All(&rtos)
	if err != nil {
		fmt.Printf("sql find info failed! %s\n", err)
	}
	p := PackInfo{}
	p.Code = 200
	p.Msg = "查询成功"
	p.Length = cnt
	p.Data = rtos
	return p
}
