/*
 * Revision History:
 *     Initial: 2018/05/05        Tong Yuehong
 */

package controllers

import (
	"encoding/json"
	"strings"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/360EntSecGroup-Skylar/excelize"

	"github.com/tongyuehong1/design-back-end/back-end/common"
	"github.com/tongyuehong1/design-back-end/back-end/models"
	"github.com/tongyuehong1/design-back-end/back-end/utility"
	"github.com/tongyuehong1/golang-project/libs/logger"
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
			fmt.Println(stu)
		}
		fmt.Println(stu)
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
			Class string `json:"class"`
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
		fmt.Println("sssss", class.Class)
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
			Class string `json:"class"`
		}
	)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &student)
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
			StuID  uint32 `json:"stuid"`
			Avatar string `json:"avatar"`
		}
	)

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &avatar)
	if err != nil {
		logger.Logger.Error("avatar Unmarshal:", err)
	} else {
		path, err := utility.SaveAvatar(avatar.StuID, avatar.Avatar)
		if err != nil {
			logger.Logger.Error("save avatar", err)
		}
		ip := "http://192.168.0.103:8080"
		path = strings.Replace(path, ".", ip, 1)
		err = models.StudentServer.UpAvatar(avatar.StuID, path)
		if err != nil {
			logger.Logger.Error("models", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlNotFound}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}
