package routers

import (
	"github.com/tongyuehong1/design-back-end/back-end/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("user/register",&controllers.UserController{}, "post:Register")
	beego.Router("user/login",&controllers.UserController{}, "post:Login")

	beego.Router("student/insert",&controllers.StudentController{}, "post:Insert")
	beego.Router("student/change",&controllers.StudentController{}, "post:Modify")
	beego.Router("student/getall",&controllers.StudentController{}, "post:GetAll")
	beego.Router("student/getleader",&controllers.StudentController{}, "post:GetLeaders")
	beego.Router("student/getone",&controllers.StudentController{}, "post:GetOne")

	beego.Router("teacher/add",&controllers.TeacherController{}, "Post:AddTeacher")
	beego.Router("teacher/change",&controllers.TeacherController{}, "Post:ChangeTech")
	beego.Router("teacher/get",&controllers.TeacherController{}, "Post:GetTeacher")
}
