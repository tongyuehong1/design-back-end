/*
 * Revision History:
 *     Initial: 2018/06/06        Tong Yuehong
 */

package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego"

	"github.com/tongyuehong1/design-back-end/back-end/common"
	"github.com/tongyuehong1/design-back-end/back-end/models"
	"github.com/tongyuehong1/design-back-end/back-end/utility"
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
		info  GradeFile
	)

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &info)
	if err != nil {
		logger.Logger.Error("add grade Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
		this.ServeJSON()
	}

	filename, err := utility.SaveGrade(info.Class, info.Subject, info.Grade)
	xlsx, err := excelize.OpenFile(filename)
	if err != nil {
		logger.Logger.Error("open grade excel error: ", err)
	}

	rows := xlsx.GetRows("Sheet1")
	for i, row := range rows {
		if i != 0 {
			grade.Name = row[1]
			grade.Class = row[0]
			grade.Subject = row[2]
			grade.Grade = row[3]
			fmt.Println(grade)
		}
		err = models.GradeServer.Insert(grade)
		if err != nil {
			logger.Logger.Error("insert student error:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		}
	}

	this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}

	this.ServeJSON()
}

// GetOne -
func GetOne(this *GradeController) {
	var (
		student struct {
			Name  string `json:"name"`
			Class string `json:"class"`
		}
	)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &student)
	if err != nil {
		logger.Logger.Error("get one's grade Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		grade, err := models.GradeServer.GetOne(student.Name, student.Class)
		if err != nil {
			logger.Logger.Error("get one's grade model", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: grade}
		}
	}

	this.ServeJSON()
}

func GetAll(this *GradeController) {
	var (
		class struct {
			Class string `json:"class"`
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
