/*
 * MIT License
 *
 * Copyright (c) 2018 SmartestEE Co., Ltd..
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2018/05/05        Tong Yuehong
 */

package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/tongyuehong1/design-back-end/back-end/models"
	"github.com/tongyuehong1/golang-project/libs/logger"
	"github.com/tongyuehong1/design-back-end/back-end/common"
)

type StudentController struct {
	beego.Controller
}

func (this *StudentController) Insert() {
	student := models.Student{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &student)
	if err != nil {
		logger.Logger.Error("add student Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.StudentServer.Insert(student)
		if err != nil {
			logger.Logger.Error("Insert student", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}

func (this *StudentController) Modify() {
	info := models.Student{}
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

func (this *StudentController) GetLeaders() {
	var (
		class string
	)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &class)
	if err != nil {
		logger.Logger.Error("change student info Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		students, err := models.StudentServer.GetLeaders(class)
		if err != nil {
			logger.Logger.Error("change student info", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: students}
		}
	}

	this.ServeJSON()
}

func (this *StudentController) GetAll() {
	var (
		class string
	)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &class)
	if err != nil {
		logger.Logger.Error("change student info Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		students, err := models.StudentServer.GetAll(class)
		if err != nil {
			logger.Logger.Error("change student info", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: students}
		}
	}

	this.ServeJSON()
}
