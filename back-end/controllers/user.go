/*
 * Revision History:
 *     Initial: 2018/05/06        Tong Yuehong
 */

package controllers

import (
	"encoding/json"
	"errors"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/tongyuehong1/design-back-end/back-end/common"
	"github.com/tongyuehong1/design-back-end/back-end/models"
	"github.com/tongyuehong1/golang-project/libs/logger"
	"fmt"
)

// UserController -
type UserController struct {
	beego.Controller
}

var InvalidObjectId = errors.New("invalid input to ObjectIdHex: ")

// Register -
func (this *UserController) Register() {
	user := models.User{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)

	if err != nil {
		logger.Logger.Error("register Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		if user.Class == "" || user.Name ==""|| user.PassWord== "" {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
			logger.Logger.Error("complete the information", InvalidObjectId)
			this.ServeJSON()
		} else {
			err := models.UserServer.Register(user)
			if err != nil {
				logger.Logger.Error("register", err)
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			} else {
				if user.Role == "teacher" {
					err := models.TeacherServer.AddTeacher(user.Name, user.Class)
					if err != nil {
						logger.Logger.Error("add teacher", err)
						this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
					}
					this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
				} else {
					this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
				}
			}
		}
	}

	this.ServeJSON()
}

// Login -
func (this *UserController) Login() {
	var (
		login struct {
			Name  string `json:"name"`
			Class string `json:"className"`
			Pass  string `json:"pass"`
			Role  string `json:"role"`
		}
	)

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &login)
	if err != nil {
		logger.Logger.Error("login Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		flag, class, err := models.UserServer.Login(login.Name, login.Class, login.Pass, login.Role)
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

// Drop -
func (this *StudentController) Drop() {
	var (
		user struct {
			Name  string `json:"name"`
			Class string `json:"className"`
		}
	)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	fmt.Println("aaaa", user)
	if err != nil {
		logger.Logger.Error("delete user info Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.UserServer.Drop(user.Name, user.Class)
		if err != nil {
			logger.Logger.Error("delete user info", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}

	this.ServeJSON()
}

// classes -
func (this *UserController) Classes() {
		class, err := models.UserServer.GetClasses()
		if err != nil {
			logger.Logger.Error("get class", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: class}
		}

	this.ServeJSON()
}
