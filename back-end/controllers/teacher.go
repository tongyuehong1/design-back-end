/*
 * Revision History:
 *     Initial: 2018/05/06        Tong Yuehong
 */

package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/tongyuehong1/design-back-end/back-end/common"
	"github.com/tongyuehong1/design-back-end/back-end/models"
	"github.com/tongyuehong1/golang-project/libs/logger"
	"fmt"
)

// TeacherController -
type TeacherController struct {
	beego.Controller
}

// AddTeacher -
func (this *TeacherController) AddTeacher() {
	teacher := models.Teacher{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &teacher)
	if err != nil {
		logger.Logger.Error("add teacher Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.TeacherServer.AddTeacher(teacher)
		if err != nil {
			logger.Logger.Error("Add teacher", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
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
		err := models.TeacherServer.ChangeTeacher(teacher)
		if err != nil {
			logger.Logger.Error("change teacher info", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
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
