/*
 * Revision History:
 *     Initial: 2018/05/06        Tong Yuehong
 */

package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"fmt"

	"github.com/tongyuehong1/design-back-end/back-end/common"
	"github.com/tongyuehong1/design-back-end/back-end/models"
	"github.com/tongyuehong1/golang-project/libs/logger"
)

// UserController -
type UserController struct {
	beego.Controller
}

// Register -
func (this *UserController) Register() {
	user := models.User{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err != nil {
		logger.Logger.Error("register Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.UserServer.Register(user)
		if err != nil {
			logger.Logger.Error("register", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}

// Login -
func (this *UserController) Login() {
	var (
		login struct {
			Name  string `json:"name"`
			Class string `json:"specialities"`
			Pass  string `json:"pass"`
		}
	)

	fmt.Println(string(this.Ctx.Input.RequestBody))
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &login)
	if err != nil {
		logger.Logger.Error("login Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		fmt.Println("fgg", login)
		flag, class, err := models.UserServer.Login(login.Name, login.Class, login.Pass)
		if err != nil {
			if err == orm.ErrNoRows {
				logger.Logger.Error("noadmin：", err)
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidUser}
			} else {
				logger.Logger.Error("models： ", err)

				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			}
		} else {
			if !flag {
				logger.Logger.Debug("Wrong Pass!")
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrWrongPass}
			} else {
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: class}
			}
		}
	}
	this.ServeJSON()
}
