/*
 * Revision History:
 *     Initial: 2018/05/05        Tong Yuehong
 */

package controllers

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"

	"github.com/tongyuehong1/design-back-end/back-end/common"
	"github.com/tongyuehong1/design-back-end/back-end/models"
	"github.com/tongyuehong1/design-back-end/back-end/utility"
	"github.com/tongyuehong1/golang-project/libs/logger"
)

// StudentController -
type StudentController struct {
	beego.Controller
}

// Insert -
func (this *StudentController) Insert() {
	student := models.Student{}

	var info struct {
		UserID uint32 `json:"userid"`
	}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &info)
	if err != nil {
		logger.Logger.Error("add student Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err = models.StudentServer.Insert(student)
		if err != nil {
			logger.Logger.Error("Insert student", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

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
			Class string `json:"class"`
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
