/*
 * Revision History:
 *     Initial: 2018/05/13        Tong Yuehong
 */

package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/tongyuehong1/design-back-end/back-end/common"
	"github.com/tongyuehong1/design-back-end/back-end/models"
	"github.com/tongyuehong1/golang-project/libs/logger"
)

// MessageController -
type MessageController struct {
	beego.Controller
}

// Publish -
func (this *MessageController) Publish() {
	message := models.Message{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &message)
	if err != nil {
		logger.Logger.Error("add student Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.MessageServer.Publish(message)
		if err != nil {
			logger.Logger.Error("Insert message", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}

// Delete -
func (this *MessageController) Delete() {
	var id struct {
		ID int8 `json:"id"`
	}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &id)
	if err != nil {
		logger.Logger.Error("add student Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.MessageServer.Delete(id.ID)
		if err != nil {
			logger.Logger.Error("Insert student", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}

// Show -
func (this *MessageController) Show() {
	var class struct {
		Class string `json:"class"`
	}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &class)
	if err != nil {
		logger.Logger.Error("show message Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		message, err := models.MessageServer.Show(class.Class)
		if err != nil {
			logger.Logger.Error("Show message", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: message}
		}
	}

	this.ServeJSON()
}
