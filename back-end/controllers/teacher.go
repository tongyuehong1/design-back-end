/*
 * Revision History:
 *     Initial: 2018/05/06        Tong Yuehong
 */

package controllers

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"

	"github.com/tongyuehong1/design-back-end/back-end/common"
	"github.com/tongyuehong1/design-back-end/back-end/models"
	"github.com/tongyuehong1/golang-project/libs/logger"
	"github.com/tongyuehong1/design-back-end/back-end/utility"
	"fmt"
)

// TeacherController -
type TeacherController struct {
	beego.Controller
}

// ChangeTech -
func (this *TeacherController) ChangeTech() {
	teacher := models.Teacher{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &teacher)
	fmt.Println("1111", teacher)
	if err != nil {
		logger.Logger.Error("change teacher's info Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		if teacher.Phone == "" || teacher.Name == ""|| teacher.Office== "" || teacher.Sex=="" {
			logger.Logger.Error("change teacher's info Unmarshal:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
		}else {
			err := models.TeacherServer.ChangeTeacher(teacher)
			if err != nil {
				logger.Logger.Error("change teacher info", err)
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			} else {
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
			}
		}
	}

	this.ServeJSON()
}

// GetTeacher -
func (this *TeacherController) GetTeacher() {
	var class struct {
		Class string `json:"className"`
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &class)
	if err != nil {
		logger.Logger.Error("change student info Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		student, err := models.TeacherServer.GetOne(class.Class)
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
func (this *StudentController) TechAvatar() {
	var (
		avatar struct {
			Name  string `json:"name"`
			Avatar string `json:"avatar"`
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
			ip := "http://10.0.0.43:21001"
			path = strings.Replace(path, ".", ip, 1)
			err = models.TeacherServer.TechAvatar(avatar.Name, path)
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
