/*
 * Revision History:
 *     Initial: 2018/05/05        Tong Yuehong
 */

package controllers

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/360EntSecGroup-Skylar/excelize"

	"github.com/tongyuehong1/design-back-end/back-end/common"
	"github.com/tongyuehong1/design-back-end/back-end/models"
	"github.com/tongyuehong1/design-back-end/back-end/utility"
	"github.com/tongyuehong1/golang-project/libs/logger"
	"fmt"
)

// StudentController -
type StudentController struct {
	beego.Controller
}

type User struct {
	Id uint32 `json:"id"`
}

type Info struct {
	Class string `json:"class"`
	File  string `json:"name"`
}

type student struct {
		Class     string `orm:"column(class)"     json:"className"`
		Name      string `orm:"column(name)"      json:"name"`
		StudentID string `orm:"column(studentid)" json:"studentid"`
		Sex       string `orm:"column(sex)"       json:"sex"`
		Age       string `orm:"column(age)"       json:"age"`
		Phone     string `orm:"column(phone)"     json:"phone"`
		Address   string `orm:"column(address)"   json:"address"`
		Duty      string `orm:"column(duty)"      json:"duty"`
	}


// Add -
func (this *StudentController) Add() {
	var stu student
	var one models.Student
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &stu)
	fmt.Println("1111", stu)
	if err != nil {
		logger.Logger.Error("add student info Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		if stu.Sex == ""|| stu.Age== ""||stu.Duty==""||stu.Class==""||stu.Phone==""||stu.Address==""||stu.StudentID=="" ||stu.Phone=="" {
			logger.Logger.Error("add student info Unmarshal:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
		}
		one.Name = stu.Name
		one.Sex = stu.Sex
		one.Class = stu.Class
		one.Phone = stu.Phone
		one.Address = stu.Address
		one.StudentID = stu.StudentID
		one.Duty = stu.Duty
		one.Age = stu.Age
		err := models.StudentServer.Insert(one)
		if err != nil {
			logger.Logger.Error("add student info", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}
// Insert -
func (this *StudentController) Insert() {
	var (
		stu models.Student
		info Info
	)

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &info)
	if err != nil {
		logger.Logger.Error("add student Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
		this.ServeJSON()
	}

	filename, err := utility.SaveFile(info.Class, info.File)
	xlsx, err := excelize.OpenFile(filename)
	if err != nil {
		logger.Logger.Error("open student excel error: ", err)
	}

	rows := xlsx.GetRows("Sheet1")
	for i, row := range rows {
		if i != 0 {
			stu.Name = row[0]
			stu.Class = row[1]
			stu.StudentID = row[2]
			stu.Sex = row[3]
			stu.Age = row[4]
			stu.Phone = row[5]
			stu.Duty = row[6]
			stu.Isonly = row[7]
			stu.Address = row[8]
		}
		err = models.StudentServer.Insert(stu)
		if err != nil {
			logger.Logger.Error("insert student error:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		}
	}

	this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}

	this.ServeJSON()
}

// Modify -
func (this *StudentController) Modify() {
	var info = models.Info{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &info)
	if err != nil {
		logger.Logger.Error("change student info Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.StudentServer.ModifyStudent(info)
		if err != nil {
			logger.Logger.Error("change student info", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}

// GetLeaders -
func (this *StudentController) GetLeaders() {
	var (
		class struct {
			Class string `json:"className"`
		}
	)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &class)
	if err != nil {
		logger.Logger.Error("change student info Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		students, err := models.StudentServer.GetLeaders(class.Class)
		if err != nil {
			logger.Logger.Error("change student info", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: students}
		}
	}

	this.ServeJSON()
}

// GetAll -
func (this *StudentController) GetAll() {
	var (
		class struct {
			Class string `json:"className"`
		}
	)

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &class)
	if err != nil {
		logger.Logger.Error("change student info Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		students, err := models.StudentServer.GetAll(class.Class)
		if err != nil {
			logger.Logger.Error("change student info", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: students}
		}
	}

	this.ServeJSON()
}

// GetOne -
func (this *StudentController) GetOne() {
	var (
		student struct {
			Name  string `json:"name"`
			Class string `json:"className"`
		}
	)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &student)
	fmt.Println("aaaa", student)
	if err != nil {
		logger.Logger.Error("change student info Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		student, err := models.StudentServer.GetOne(student.Name, student.Class)
		if err != nil {
			logger.Logger.Error("change student info", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: student}
		}
	}

	this.ServeJSON()
}

// UpAvatar -
func (this *StudentController) UpAvatar() {
	var (
		avatar struct {
			Name  string `json:"name"`
			Avatar string `json:"avatar"`
			Class  string `json:"className"`
		}
	)

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &avatar)
	if err != nil {
		logger.Logger.Error("avatar Unmarshal:", err)
	} else {
		path, err := utility.SaveAvatar(avatar.Name, avatar.Avatar)
		if err != nil {
			logger.Logger.Error("save avatar", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrNotFound}
		} else {
			ip := "http://192.168.43.218:21001"
			path = strings.Replace(path, ".", ip, 1)
			err = models.StudentServer.UpAvatar(avatar.Name, path,avatar.Class)
			if err != nil {
				logger.Logger.Error("models", err)
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlNotFound}
			} else {
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: path}
			}
		}
	}

	this.ServeJSON()
}

// Delete
func (this *StudentController) Delete() {
	var (
		student struct {
			Name  string `json:"name"`
			Class string `json:"className"`
		}
	)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &student)
	fmt.Println("aaaa", student)
	if err != nil {
		logger.Logger.Error("delete student info Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.StudentServer.Delete(student.Name, student.Class)
		if err != nil {
			logger.Logger.Error("delete student info", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			err := models.UserServer.Drop(student.Name, student.Class)
			if err != nil {
				logger.Logger.Error("delete user info", err)
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			}
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}