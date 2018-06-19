/*
 * Revision History:
 *     Initial: 2018/06/06        Tong Yuehong
 */

package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"

	"github.com/tongyuehong1/design-back-end/back-end/common"
	"github.com/tongyuehong1/design-back-end/back-end/models"
	"github.com/tongyuehong1/golang-project/libs/logger"
)

type GradeController struct {
	beego.Controller
}

type GradeFile struct {
	Class   string `json:"class"`
	Subject string `json:"project"`
	Grade   string `json:"grade"`
}

// AddGrade -
func (this *GradeController) AddGrade() {
	var (
		grade models.Grade
	)

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &grade)
	fmt.Println("2222", grade)
	if err != nil {
		logger.Logger.Error("change student info Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.GradeServer.Insert(grade)
		if err != nil {
			logger.Logger.Error("change student info", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}

// GetOne -
func (this *GradeController) GetOne() {
	var (
		student struct {
			Name  string `json:"name"`
			Class string `json:"className"`
		}
	)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &student)
	fmt.Println("rrrrr", student)
	if err != nil {
		logger.Logger.Error("get one's grade Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		grade, err := models.GradeServer.GetOne(student.Class, student.Name)
		if err != nil {
			logger.Logger.Error("get one's grade model", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: grade}
		}
	}

	this.ServeJSON()
}

func (this *GradeController) GetAll() {
	var (
		class struct {
			Class string `json:"className"`
		}
	)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &class)
	if err != nil {
		logger.Logger.Error("get all grade Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		grades, err := models.GradeServer.GetAll(class.Class)
		if err != nil {
			logger.Logger.Error("get all grade model", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: grades}
		}
	}

	this.ServeJSON()
}
